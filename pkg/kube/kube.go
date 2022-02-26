package kube

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
	"os"

	scheme "github.com/bhojpur/application/pkg/client/clientset/versioned"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// Clients returns new Kubernetes and Bhojpur Application clients
func Clients() (*kubernetes.Clientset, scheme.Interface, error) {
	var config *rest.Config
	var err error
	pathToKubeConfig := os.Getenv("APP_DASHBOARD_KUBECONFIG")
	if pathToKubeConfig != "" {
		config, err = clientcmd.BuildConfigFromFlags("", pathToKubeConfig)
		if err != nil {
			return nil, nil, err
		}
	}

	if config == nil {
		config, err = rest.InClusterConfig()
		if err != nil {
			return nil, nil, err
		}
	}

	kubeClient, err := kubernetes.NewForConfig(config)
	if err != nil {
		return kubeClient, nil, err
	}

	appClient, err := scheme.NewForConfig(config)
	return kubeClient, appClient, err
}
