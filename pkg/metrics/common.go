package metrics

import (
	"fmt"

	"observability/pkg/core"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	responseCodeCount = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "http_response_code_counter",
		Help: "The counter which represents number of http responses",
	}, []string{"method", "status_code"})

	responseTimeHist = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "http_response_time_seconds",
			Help: "Response time of the API in seconds",
		},
		[]string{"path"},
	)
)

func RecordResponse(ctx *core.ApiContext) {
	responseCodeCount.With(prometheus.Labels{
		"status_code": fmt.Sprintf("%d", ctx.StatusCode()),
		"method":      ctx.Method(),
	}).Inc()

	responseTimeHist.With(prometheus.Labels{
		"path": ctx.Path(),
	}).Observe(ctx.ElapsedTime().Seconds())
}
