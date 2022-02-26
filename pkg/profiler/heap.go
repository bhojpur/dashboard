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
	// VHeap is the name of HeapViewer
	VHeap = "heap"
)

// HeapViewer collects the heap-stats metrics via `runtime.ReadMemStats()`
type HeapViewer struct {
	smgr  *StatsMgr
	graph *charts.Line
}

// NewHeapViewer returns the HeapViewer instance
// Series: Alloc / Inuse / Sys / Idle
func NewHeapViewer() Viewer {
	graph := newBasicView(VHeap)
	graph.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "Heap Memory"}),
		charts.WithYAxisOpts(opts.YAxis{Name: "Size", AxisLabel: &opts.AxisLabel{Formatter: "{value} MB"}}),
	)
	graph.AddSeries("Alloc", []opts.LineData{}).
		AddSeries("Inuse", []opts.LineData{}).
		AddSeries("Sys", []opts.LineData{}).
		AddSeries("Idle", []opts.LineData{})

	return &HeapViewer{graph: graph}
}
func (vr *HeapViewer) SetStatsMgr(smgr *StatsMgr) {
	vr.smgr = smgr
}
func (vr *HeapViewer) Name() string {
	return VHeap
}

func (vr *HeapViewer) View() *charts.Line {
	return vr.graph
}

func (vr *HeapViewer) Serve(w http.ResponseWriter, _ *http.Request) {
	vr.smgr.Tick()

	metrics := Metrics{
		Values: []float64{
			fixedPrecision(float64(memstats.Stats.HeapAlloc)/1024/1024, 2),
			fixedPrecision(float64(memstats.Stats.HeapInuse)/1024/1024, 2),
			fixedPrecision(float64(memstats.Stats.HeapSys)/1024/1024, 2),
			fixedPrecision(float64(memstats.Stats.HeapIdle)/1024/1024, 2),
		},
		Time: memstats.T,
	}

	bs, _ := json.Marshal(metrics)
	w.Write(bs)
}
