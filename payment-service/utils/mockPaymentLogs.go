package utils

import (
	"time"

	"github.com/pnaskardev/EventLens/payment-service/pkg/entities"
)

var PaymentLogs = []entities.PaymentLog{
	{
		Timestamp: time.Now().Add(-10 * time.Minute),
		Level:     "INFO",
		TraceID:   "trace_2001",
		OrderID:   "order_5001",
		Message:   "Payment initiated",
		Meta: struct {
			Amount float64 `json:"amount"`
			Items  int64   `json:"items"`
		}{
			Amount: 1499.99,
			Items:  2,
		},
	},
	{
		Timestamp: time.Now().Add(-9 * time.Minute),
		Level:     "INFO",
		TraceID:   "trace_2001",
		OrderID:   "order_5001",
		Message:   "Payment successful",
		Meta: struct {
			Amount float64 `json:"amount"`
			Items  int64   `json:"items"`
		}{
			Amount: 1499.99,
			Items:  2,
		},
	},
	{
		Timestamp: time.Now().Add(-8 * time.Minute),
		Level:     "INFO",
		TraceID:   "trace_2002",
		OrderID:   "order_5002",
		Message:   "Payment initiated",
		Meta: struct {
			Amount float64 `json:"amount"`
			Items  int64   `json:"items"`
		}{
			Amount: 799.50,
			Items:  1,
		},
	},
	{
		Timestamp: time.Now().Add(-7 * time.Minute),
		Level:     "ERROR",
		TraceID:   "trace_2002",
		OrderID:   "order_5002",
		Message:   "Payment failed",
		Meta: struct {
			Amount float64 `json:"amount"`
			Items  int64   `json:"items"`
		}{
			Amount: 799.50,
			Items:  1,
		},
	},
	{
		Timestamp: time.Now().Add(-6 * time.Minute),
		Level:     "WARN",
		TraceID:   "trace_2002",
		OrderID:   "order_5002",
		Message:   "Retrying payment",
		Meta: struct {
			Amount float64 `json:"amount"`
			Items  int64   `json:"items"`
		}{
			Amount: 799.50,
			Items:  1,
		},
	},
	{
		Timestamp: time.Now().Add(-5 * time.Minute),
		Level:     "INFO",
		TraceID:   "trace_2002",
		OrderID:   "order_5002",
		Message:   "Payment successful after retry",
		Meta: struct {
			Amount float64 `json:"amount"`
			Items  int64   `json:"items"`
		}{
			Amount: 799.50,
			Items:  1,
		},
	},
	{
		Timestamp: time.Now().Add(-4 * time.Minute),
		Level:     "INFO",
		TraceID:   "trace_2003",
		OrderID:   "order_5003",
		Message:   "Payment initiated",
		Meta: struct {
			Amount float64 `json:"amount"`
			Items  int64   `json:"items"`
		}{
			Amount: 2499.00,
			Items:  4,
		},
	},
	{
		Timestamp: time.Now().Add(-3 * time.Minute),
		Level:     "ERROR",
		TraceID:   "trace_2003",
		OrderID:   "order_5003",
		Message:   "Payment failed due to gateway timeout",
		Meta: struct {
			Amount float64 `json:"amount"`
			Items  int64   `json:"items"`
		}{
			Amount: 2499.00,
			Items:  4,
		},
	},
	{
		Timestamp: time.Now().Add(-2 * time.Minute),
		Level:     "ERROR",
		TraceID:   "trace_2004",
		OrderID:   "order_5004",
		Message:   "Payment failed due to fraud detection",
		Meta: struct {
			Amount float64 `json:"amount"`
			Items  int64   `json:"items"`
		}{
			Amount: 599.99,
			Items:  1,
		},
	},
	{
		Timestamp: time.Now().Add(-1 * time.Minute),
		Level:     "INFO",
		TraceID:   "trace_2005",
		OrderID:   "order_5005",
		Message:   "Payment processing delayed",
		Meta: struct {
			Amount float64 `json:"amount"`
			Items  int64   `json:"items"`
		}{
			Amount: 1299.75,
			Items:  3,
		},
	},
}
