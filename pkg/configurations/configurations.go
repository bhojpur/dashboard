package configurations

// Copyright (c) 2018 Bhojpur Consulting Private Limited, India. All rights reserved.

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

import (
	"log"
	"os"
	"path/filepath"
	"strconv"

	scheme "github.com/bhojpur/application/pkg/client/clientset/versioned"
	"github.com/bhojpur/application/pkg/config"
	"github.com/bhojpur/application/pkg/standalone"
	"github.com/bhojpur/dashboard/pkg/age"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Configurations is an interface to interact with Bhojpur Application configurations
type Configurations interface {
	Supported() bool
	GetConfiguration(scope string, name string) Configuration
	GetConfigurations(scope string) []Configuration
}

type configurations struct {
	platform            string
	appClient           scheme.Interface
	getConfigurationsFn func(scope string) []Configuration
}

// NewConfigurations returns a new Configurations instance
func NewConfigurations(platform string, appClient scheme.Interface) Configurations {
	c := configurations{}
	c.platform = platform

	if platform == "kubernetes" {
		c.getConfigurationsFn = c.getKubernetesConfigurations
		c.appClient = appClient
	} else if platform == "standalone" {
		c.getConfigurationsFn = c.getStandaloneConfigurations
	}
	return &c
}

// Configuration represents a Bhojpur Application configuration
type Configuration struct {
	Name            string      `json:"name"`
	Kind            string      `json:"kind"`
	Created         string      `json:"created"`
	Age             string      `json:"age"`
	TracingEnabled  bool        `json:"tracingEnabled"`
	SamplingRate    string      `json:"samplingRate"`
	MetricsEnabled  bool        `json:"metricsEnabled"`
	MTLSEnabled     bool        `json:"mtlsEnabled"`
	WorkloadCertTTL string      `json:"mtlsWorkloadTTL"`
	ClockSkew       string      `json:"mtlsClockSkew"`
	Manifest        interface{} `json:"manifest"`
}

// Supported checks whether or not the supplied platform is able to access
// Bhojpur Application configurations
func (c *configurations) Supported() bool {
	return c.platform == "kubernetes" || c.platform == "standalone"
}

// GetConfiguration returns the Bhojpur Application configuration specified by name
func (c *configurations) GetConfiguration(scope string, name string) Configuration {
	confs := c.getConfigurationsFn(scope)
	for _, conf := range confs {
		if conf.Name == name {
			return conf
		}
	}
	return Configuration{}
}

// GetConfigurations returns the result of the correct platform's getConfigurations function
func (c *configurations) GetConfigurations(scope string) []Configuration {
	return c.getConfigurationsFn(scope)
}

// getKubernetesConfigurations returns the list of all Bhojpur Application
// Configurations in a Kubernetes cluster
func (c *configurations) getKubernetesConfigurations(scope string) []Configuration {
	confs, err := c.appClient.ConfigurationV1alpha1().Configurations(scope).List(meta_v1.ListOptions{})
	if err != nil {
		log.Println(err)
		return []Configuration{}
	}

	out := []Configuration{}
	for _, comp := range confs.Items {
		newConfiguration := Configuration{
			Name:            comp.Name,
			Kind:            comp.Kind,
			Created:         comp.CreationTimestamp.Format("2006-01-02 15:04.05"),
			Age:             age.GetAge(comp.CreationTimestamp.Time),
			TracingEnabled:  tracingEnabled(comp.Spec.TracingSpec.SamplingRate),
			SamplingRate:    comp.Spec.TracingSpec.SamplingRate,
			MetricsEnabled:  comp.Spec.MetricSpec.Enabled,
			MTLSEnabled:     comp.Spec.MTLSSpec.Enabled,
			WorkloadCertTTL: comp.Spec.MTLSSpec.WorkloadCertTTL,
			ClockSkew:       comp.Spec.MTLSSpec.AllowedClockSkew,
			Manifest:        comp,
		}
		out = append(out, newConfiguration)
	}
	return out
}

// getStandaloneConfigurations returns the list of Bhojpur Application Configurations Statuses
func (c *configurations) getStandaloneConfigurations(scope string) []Configuration {
	configurationsDirectory := filepath.Dir(standalone.DefaultConfigFilePath())
	standaloneConfigurations := []Configuration{}
	err := filepath.Walk(configurationsDirectory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Printf("Failure accessing path %s: %v\n", path, err)
			return err
		}
		if info.IsDir() && info.Name() != filepath.Base(configurationsDirectory) {
			return filepath.SkipDir
		} else if !info.IsDir() && filepath.Ext(path) == ".yaml" {
			comp, content, err := config.LoadStandaloneConfiguration(path)
			if err != nil {
				log.Printf("Failure reading configuration file %s: %v\n", path, err)
				return err
			}

			newConfiguration := Configuration{
				Name:            comp.Name,
				Kind:            comp.Kind,
				Created:         info.ModTime().Format("2006-01-02 15:04.05"),
				Age:             age.GetAge(info.ModTime()),
				TracingEnabled:  tracingEnabled(comp.Spec.TracingSpec.SamplingRate),
				SamplingRate:    comp.Spec.TracingSpec.SamplingRate,
				MetricsEnabled:  comp.Spec.MetricSpec.Enabled,
				MTLSEnabled:     comp.Spec.MTLSSpec.Enabled,
				WorkloadCertTTL: comp.Spec.MTLSSpec.WorkloadCertTTL,
				ClockSkew:       comp.Spec.MTLSSpec.AllowedClockSkew,
				Manifest:        content,
			}

			if newConfiguration.Kind == "Configuration" {
				standaloneConfigurations = append(standaloneConfigurations, newConfiguration)
			}
			return nil
		}
		return nil
	})
	if err != nil {
		log.Printf("error walking the path %q: %v\n", configurationsDirectory, err)
		return []Configuration{}
	}
	return standaloneConfigurations
}

// tracingEnabled checks if tracing is enabled for a configuration
func tracingEnabled(samplingRate string) bool {
	sr, err := strconv.ParseFloat(samplingRate, 32)
	if err != nil {
		return false
	}
	return sr > 0
}
