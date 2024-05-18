package app_errors

import "errors"

var (
	ErrOrderNotFound    = errors.New("order not found")
	ErrCantChangeStatus = errors.New("can't change order status")
	ErrCantCancelOrder  = errors.New("can't cancel order")
)
