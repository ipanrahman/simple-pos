package http

import "github.com/ipan97/simple-pos/internal/core"

type OrderHTTPHandler struct {
	core.BaseHandler
}

func NewOrderHTTPHandler() *OrderHTTPHandler {
	return &OrderHTTPHandler{}
}
