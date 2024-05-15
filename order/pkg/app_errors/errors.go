package app_errors

import "errors"

var (
	OrderCancelledError = errors.New("order already cancelled")
	OrderNotFoundError  = errors.New("order not found")
)
