package core

import (
	"encoding/json"
	"net/http"
	"time"
)

type ApiContext struct {
	w          http.ResponseWriter
	r          *http.Request
	timeStart  time.Time
	statusCode int
}

func NewApiContext(w http.ResponseWriter, r *http.Request) *ApiContext {
	return &ApiContext{
		w:         w,
		r:         r,
		timeStart: time.Now(),
	}
}

func (ctx *ApiContext) JSON(v any, statusCode int) error {
	payload, err := json.Marshal(v)
	if err != nil {
		return err
	}

	ctx.statusCode = statusCode
	ctx.w.WriteHeader(statusCode)
	_, err = ctx.w.Write(payload)

	return err
}

func (ctx *ApiContext) Method() string {
	return ctx.r.Method
}

func (ctx *ApiContext) Path() string {
	return ctx.r.URL.Path
}

func (ctx *ApiContext) StatusCode() int {
	return ctx.statusCode
}

func (ctx *ApiContext) ElapsedTime() time.Duration {
	return time.Since(ctx.timeStart)
}
