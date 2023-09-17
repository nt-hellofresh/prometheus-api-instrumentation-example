package pkg

import (
	"errors"
	"math/rand"
	"net/http"

	"observability/pkg/core"
	"observability/pkg/metrics"
)

func homeRoute(ctx *core.ApiContext) error {
	resp := testResponse{
		Message: "Hello, World!",
	}
	return ctx.JSON(resp, http.StatusOK)
}

func headsOrTailsRoute(ctx *core.ApiContext) error {
	outcome := flipCoin()
	metrics.IncrementHeadsOrTailsMetrics(outcome)
	return ctx.JSON(map[string]string{"outcome": outcome}, http.StatusOK)
}

func routeWithError(_ *core.ApiContext) error {
	return errors.New("something went wrong")
}

func flipCoin() string {
	results := rand.Intn(2)

	if results == 0 {
		return "heads"
	}
	return "tails"
}
