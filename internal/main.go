package main

import (
	"context"
	"fmt"
	"net/http"
	"net/http/pprof"
	"time"

	"github.com/bhojpur/charts/pkg/components"
	"github.com/bhojpur/charts/pkg/templates"
	"github.com/rs/cors"

	profiler "github.com/bhojpur/dashboard/pkg/profiler"
	statics "github.com/bhojpur/dashboard/statics"
)

// ViewManager
type ViewManager struct {
	srv *http.Server

	Smgr   *profiler.StatsMgr
	Ctx    context.Context
	Cancel context.CancelFunc
	Views  []profiler.Viewer
}

// Register registers views to the ViewManager
func (vm *ViewManager) Register(views ...profiler.Viewer) {
	fmt.Sprintf("Registering %s", profiler.LinkAddr())
	vm.Views = append(vm.Views, views...)

}

// Start runs an HTTP server and begin to collect metrics
func (vm *ViewManager) Start() error {
	fmt.Println("Bhojpur Dashboard - Web Server started at", vm.srv.Addr)
	return vm.srv.ListenAndServe()
}

// Stop shutdown the http server gracefully
func (vm *ViewManager) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	vm.srv.Shutdown(ctx)
	fmt.Println("Bhojpur Dashboard - Web Server shutdown", vm.srv.Addr)
	vm.Cancel()
}

func init() {
	templates.PageTpl = `
{{- define "page" }}
<!DOCTYPE html>
<html>
    {{- template "header" . }}
<body>
<img src="https://static.bhojpur.net/image/logo.png"/>
<hr size="1px">
<style> .box { justify-content:center; display:flex; flex-wrap:wrap } </style>
<div class="box"> {{- range .Charts }} {{ template "base" . }} {{- end }} </div>
{{- template "footer" . }}
</body>
</html>
{{ end }}
`
}

// New creates a new ViewManager instance
func New() *ViewManager {
	page := components.NewPage()
	page.PageTitle = "Bhojpur Dashboard"
	page.AssetsHost = fmt.Sprintf("http://%s/profiler/statics/", profiler.LinkAddr())
	page.Assets.JSAssets.Add("jquery.min.js")

	mgr := &ViewManager{
		srv: &http.Server{
			Addr:           profiler.Addr(),
			ReadTimeout:    time.Minute,
			WriteTimeout:   time.Minute,
			MaxHeaderBytes: 1 << 20,
		},
	}
	mgr.Ctx, mgr.Cancel = context.WithCancel(context.Background())
	mgr.Register(
		profiler.NewGoroutinesViewer(),
		profiler.NewHeapViewer(),
		profiler.NewStackViewer(),
		profiler.NewGCNumViewer(),
		profiler.NewGCSizeViewer(),
		profiler.NewGCCPUFractionViewer(),
	)
	smgr := profiler.NewStatsMgr(mgr.Ctx)
	for _, v := range mgr.Views {
		v.SetStatsMgr(smgr)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)

	for _, v := range mgr.Views {
		page.AddCharts(v.View())
		mux.HandleFunc("/profiler/view/"+v.Name(), v.Serve)
	}

	mux.HandleFunc("/profiler", func(w http.ResponseWriter, _ *http.Request) {
		page.Render(w)
	})

	staticsPrev := "/profiler/statics/"
	mux.HandleFunc(staticsPrev+"echarts.min.js", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte(statics.EchartJS))
	})

	mux.HandleFunc(staticsPrev+"jquery.min.js", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte(statics.JqueryJS))
	})

	mux.HandleFunc(staticsPrev+"themes/westeros.js", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte(statics.WesterosJS))
	})

	mux.HandleFunc(staticsPrev+"themes/macarons.js", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte(statics.MacaronsJS))
	})

	mgr.srv.Handler = cors.AllowAll().Handler(mux)
	return mgr
}

func main() {
	mgr := New()

	// Start() runs an HTTP server at `localhost:18066` by default.
	go mgr.Start()

	// Stop() will shutdown the HTTP server gracefully
	// mgr.Stop()

	// busy working....
	time.Sleep(time.Minute)
}
