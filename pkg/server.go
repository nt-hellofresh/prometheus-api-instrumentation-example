package pkg

import (
	"log"
	"net/http"
	"observability/pkg/app"

	"observability/pkg/core"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func RegisterRoutes() {
	log.Printf("Registering routes...")
	http.HandleFunc("/", makeHandler(app.HomeRoute))
	http.HandleFunc("/heads-or-tails", makeHandler(app.HeadsOrTailsRoute))
	http.HandleFunc("/bad-route", makeHandler(app.RouteWithError))
	http.Handle("/metrics", promhttp.Handler())
}

func Start(port string) {
	log.Printf("start listening on port %v", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func makeHandler(fn core.HandlerFunc) http.HandlerFunc {
	handler := fn
	handler = unhandledErrorMiddleware(handler)
	handler = responseMetricsMiddleware(handler)

	return func(w http.ResponseWriter, r *http.Request) {
		ctx := core.NewApiContext(w, r)
		_ = handler(ctx)
		log.Printf("[%v] %v (%d)", ctx.Method(), ctx.Path(), ctx.StatusCode())
	}
}
