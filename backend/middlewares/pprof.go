package middlewares

import (
	"net/http/pprof"

	"github.com/bndrmrtn/go-gale"
)

func RegisterPprof(r gale.Router) {
	r = r.Group("/debug")

	r.All("/pprof", gale.Adaptor(pprof.Index))
	r.All("/allocs", gale.Adaptor(pprof.Handler("allocs").ServeHTTP))
	r.All("/block", gale.Adaptor(pprof.Handler("block").ServeHTTP))
	r.All("/cmdline", gale.Adaptor(pprof.Cmdline))
	r.All("/goroutine", gale.Adaptor(pprof.Handler("goroutine").ServeHTTP))
	r.All("/heap", gale.Adaptor(pprof.Handler("heap").ServeHTTP))
	r.All("/mutex", gale.Adaptor(pprof.Handler("mutex").ServeHTTP))
	r.All("/profile", gale.Adaptor(pprof.Profile))
	r.All("/threadcreate", gale.Adaptor(pprof.Handler("threadcreate").ServeHTTP))
	r.All("/symbol", gale.Adaptor(pprof.Symbol))
	r.All("/trace", gale.Adaptor(pprof.Trace))
}
