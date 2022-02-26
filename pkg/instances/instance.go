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

// Instance describes a Bhojpur Application sidecar instance information
type Instance struct {
	AppID            string `json:"appID"`
	HTTPPort         int    `json:"httpPort"`
	GRPCPort         int    `json:"grpcPort"`
	AppPort          int    `json:"appPort"`
	Command          string `json:"command"`
	Age              string `json:"age"`
	Created          string `json:"created"`
	PID              int    `json:"pid"`
	Replicas         int    `json:"replicas"`
	Address          string `json:"address"`
	SupportsDeletion bool   `json:"supportsDeletion"`
	SupportsLogs     bool   `json:"supportsLogs"`
	Manifest         string `json:"manifest"`
	Status           string `json:"status"`
	Labels           string `json:"labels"`
	Selector         string `json:"selector"`
	Config           string `json:"config"`
}

// StatusOutput represents the status of a named Bhojpur Application resource.
type StatusOutput struct {
	Service   string `json:"service"`
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Healthy   string `json:"healthy"`
	Status    string `json:"status"`
	Version   string `json:"version"`
	Age       string `json:"age"`
	Created   string `json:"created"`
}

// MetadataOutput represents a metadata api call response
type MetadataOutput struct {
	ID       string                      `json:"id"`
	Actors   []MetadataActiveActorsCount `json:"actors"`
	Extended map[string]interface{}      `json:"extended"`
}

// MetadataActiveActorsCount represents actor metadata: type and count
type MetadataActiveActorsCount struct {
	Type  string `json:"type"`
	Count int    `json:"count"`
}
