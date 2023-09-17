package pkg

import (
	"net/http"

	"observability/pkg/core"
	"observability/pkg/metrics"
)

func responseMetricsMiddleware(next core.Handler) core.Handler {
	return func(ctx *core.ApiContext) error {
		err := next(ctx)
		metrics.RecordResponse(ctx)
		return err
	}
}

func unhandledErrorMiddleware(next core.Handler) core.Handler {
	return func(ctx *core.ApiContext) error {
		err := next(ctx)
		if err != nil {
			_ = ctx.JSON(map[string]string{"error": err.Error()}, http.StatusInternalServerError)
		}
		return err
	}
}
