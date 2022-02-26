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
	"runtime"
	"time"

	"github.com/bhojpur/charts/pkg/charts"
	"github.com/bhojpur/charts/pkg/opts"
)

const (
	// VGoroutine is the name of GoroutinesViewer
	VGoroutine = "goroutine"
)

// GoroutinesViewer collects the goroutine number metric via `runtime.NumGoroutine()`
type GoroutinesViewer struct {
	smgr  *StatsMgr
	graph *charts.Line
}

// NewGoroutinesViewer returns the GoroutinesViewer instance
// Series: Goroutines
func NewGoroutinesViewer() Viewer {
	graph := newBasicView(VGoroutine)
	graph.SetGlobalOptions(
		charts.WithYAxisOpts(opts.YAxis{Name: "Num"}),
		charts.WithTitleOpts(opts.Title{Title: "Go Routines"}),
	)
	graph.AddSeries("Goroutines", []opts.LineData{})

	return &GoroutinesViewer{graph: graph}
}

func (vr *GoroutinesViewer) SetStatsMgr(smgr *StatsMgr) {
	vr.smgr = smgr
}

func (vr *GoroutinesViewer) Name() string {
	return VGoroutine
}

func (vr *GoroutinesViewer) View() *charts.Line {
	return vr.graph
}

func (vr *GoroutinesViewer) Serve(w http.ResponseWriter, _ *http.Request) {
	vr.smgr.Tick()

	metrics := Metrics{
		Values: []float64{float64(runtime.NumGoroutine())},
		Time:   time.Now().Format(defaultCfg.TimeFormat),
	}

	bs, _ := json.Marshal(metrics)
	w.Write(bs)
}
