package middlewares

import (
	"net/http/pprof"

	"github.com/bndrmrtn/go-gale"
)

func RegisterPprof(r gale.Router) {
	r = r.Group("/debug/pprof")

	r.All("/", gale.Adaptor(pprof.Index))
	r.All("/cmdline", gale.Adaptor(pprof.Cmdline))
	r.All("/profile", gale.Adaptor(pprof.Profile))
	r.All("/symbol", gale.Adaptor(pprof.Symbol))
	r.All("/heap", gale.Adaptor(pprof.Handler("heap").ServeHTTP))
}
