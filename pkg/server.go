package pkg

import (
	"log"
	"net/http"

	"observability/pkg/core"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func RegisterRoutes() {
	log.Printf("Registering routes...")
	http.HandleFunc("/", makeHandler(homeRoute))
	http.HandleFunc("/heads-or-tails", makeHandler(headsOrTailsRoute))
	http.HandleFunc("/bad-route", makeHandler(routeWithError))
	http.Handle("/metrics", promhttp.Handler())
}

func Start(port string) {
	log.Printf("start listening on port %v", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func makeHandler(fn core.Handler) func(http.ResponseWriter, *http.Request) {
	handler := fn
	handler = unhandledErrorMiddleware(handler)
	handler = responseMetricsMiddleware(handler)

	return func(w http.ResponseWriter, r *http.Request) {
		ctx := core.NewApiContext(w, r)
		_ = handler(ctx)
		log.Printf("[%v] %v (%d)", ctx.Method(), ctx.Path(), ctx.StatusCode())
	}
}
