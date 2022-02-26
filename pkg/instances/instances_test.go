package instances

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
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/fake"
)

func newTestSimpleK8s(objects ...runtime.Object) instances {
	client := instances{}
	client.kubeClient = fake.NewSimpleClientset(objects...)
	return client
}

func newAppControlPlanePod(name string, appName string, creationTime time.Time, state v1.ContainerState, ready bool) *v1.Pod {
	return &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:        name,
			Namespace:   "app-system",
			Annotations: map[string]string{},
			Labels: map[string]string{
				"app": appName,
			},
			CreationTimestamp: metav1.Time{
				Time: creationTime,
			},
		},
		Status: v1.PodStatus{
			ContainerStatuses: []v1.ContainerStatus{
				{
					State: state,
					Ready: ready,
				},
			},
		},
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{
					Image: name + ":0.0.1",
				},
			},
		},
	}
}

func TestControlPlaneServices(t *testing.T) {
	controlPlaneServices := []struct {
		name    string
		appName string
	}{
		{"app-operator-67d7d7bb6c-7h96c", "app-operator"},
		{"app-operator-67d7d7bb6c-2h96d", "app-operator"},
		{"app-operator-67d7d7bb6c-3h96c", "app-operator"},
		{"app-placement-server-0", "app-placement-server"},
		{"app-placement-server-1", "app-placement-server"},
		{"app-placement-server-2", "app-placement-server"},
		{"app-sentry-647759cd46-9ptks", "app-sentry"},
		{"app-sentry-647759cd46-aptks", "app-sentry"},
		{"app-sentry-647759cd46-bptks", "app-sentry"},
		{"app-sidecar-injector-74648c9dcb-5bsmn", "app-sidecar-injector"},
		{"app-sidecar-injector-74648c9dcb-6bsmn", "app-sidecar-injector"},
		{"app-sidecar-injector-74648c9dcb-7bsmn", "app-sidecar-injector"},
	}

	runtimeObj := make([]runtime.Object, len(controlPlaneServices))
	for i, s := range controlPlaneServices {
		testTime := time.Now()
		runtimeObj[i] = newAppControlPlanePod(
			s.name, s.appName,
			testTime.Add(time.Duration(-20)*time.Minute),
			v1.ContainerState{
				Running: &v1.ContainerStateRunning{
					StartedAt: metav1.Time{
						Time: testTime.Add(time.Duration(-19) * time.Minute),
					},
				},
			}, true)
	}

	k8s := newTestSimpleK8s(runtimeObj...)
	status := k8s.GetControlPlaneStatus()

	assert.Equal(t, 12, len(status), "Expected status list length to match")
}
