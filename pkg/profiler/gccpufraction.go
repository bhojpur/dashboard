package viewer

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
	"encoding/json"
	"net/http"

	"github.com/bhojpur/charts/pkg/charts"
	"github.com/bhojpur/charts/pkg/opts"
)

const (
	// VGCCPUFraction is the name of GCCPUFractionViewer
	VGCCPUFraction = "gccpufraction"
)

// GCCPUFractionViewer collects the GC-CPU fraction metric via `runtime.ReadMemStats()`
type GCCPUFractionViewer struct {
	smgr  *StatsMgr
	graph *charts.Line
}

// NewGCCPUFractionViewer returns the GCCPUFractionViewer instance
// Series: Fraction
func NewGCCPUFractionViewer() Viewer {
	graph := newBasicView(VGCCPUFraction)
	graph.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "GC CPU Fraction"}),
		charts.WithYAxisOpts(opts.YAxis{Name: "Percent", AxisLabel: &opts.AxisLabel{Formatter: "{value} %", Rotate: 35}}),
	)
	graph.AddSeries("Fraction", []opts.LineData{})

	return &GCCPUFractionViewer{graph: graph}
}

func (vr *GCCPUFractionViewer) SetStatsMgr(smgr *StatsMgr) {
	vr.smgr = smgr
}

func (vr *GCCPUFractionViewer) Name() string {
	return VGCCPUFraction
}

func (vr *GCCPUFractionViewer) View() *charts.Line {
	return vr.graph
}

func (vr *GCCPUFractionViewer) Serve(w http.ResponseWriter, _ *http.Request) {
	vr.smgr.Tick()

	metrics := Metrics{
		Values: []float64{fixedPrecision(memstats.Stats.GCCPUFraction, 6)},
		Time:   memstats.T,
	}

	bs, _ := json.Marshal(metrics)
	w.Write(bs)
}
