package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	headsCount = promauto.NewCounter(prometheus.CounterOpts{
		Name: "heads_or_tails_heads_count",
		Help: "The counter which represents the total number of heads drawn",
	})

	tailsCount = promauto.NewCounter(prometheus.CounterOpts{
		Name: "heads_or_tails_tails_count",
		Help: "The counter which represents the total number of tails drawn",
	})

	headsTailsResults = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "heads_or_tails_counter_vec",
		Help: "The counter which represents number of results drawn per outcome",
	}, []string{"outcome"})
)

func IncrementHeadsOrTailsMetrics(outcome string) {
	headsTailsResults.With(prometheus.Labels{
		"outcome": outcome,
	}).Inc()

	if outcome == "heads" {
		headsCount.Inc()
	} else {
		tailsCount.Inc()
	}
}
