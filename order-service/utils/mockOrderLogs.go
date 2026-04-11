package utils

import (
	"time"

	"github.com/pnaskardev/EventLens/order-service/pkg/entities"
)

var OrderLogs = []entities.OrderLog{
	{
		Timestamp: time.Now().Add(-12 * time.Minute),
		Level:     "INFO",
		TraceID:   "trace_1001",
		UserID:    "user_21",
		OrderID:   "order_5001",
		Message:   "Order created",
		Meta: struct {
			Amount float64 `json:"amount"`
			Items  float64 `json:"items"`
		}{
			Amount: 1499.99,
			Items:  2,
		},
	},
	{
		Timestamp: time.Now().Add(-11 * time.Minute),
		Level:     "INFO",
		TraceID:   "trace_1001",
		UserID:    "user_21",
		OrderID:   "order_5001",
		Message:   "Order validated",
		Meta: struct {
			Amount float64 `json:"amount"`
			Items  float64 `json:"items"`
		}{
			Amount: 1499.99,
			Items:  2,
		},
	},
	{
		Timestamp: time.Now().Add(-10 * time.Minute),
		Level:     "WARN",
		TraceID:   "trace_1002",
		UserID:    "user_42",
		OrderID:   "order_5002",
		Message:   "Inventory low for item",
		Meta: struct {
			Amount float64 `json:"amount"`
			Items  float64 `json:"items"`
		}{
			Amount: 799.50,
			Items:  1,
		},
	},
	{
		Timestamp: time.Now().Add(-9 * time.Minute),
		Level:     "ERROR",
		TraceID:   "trace_1002",
		UserID:    "user_42",
		OrderID:   "order_5002",
		Message:   "Order failed due to inventory shortage",
		Meta: struct {
			Amount float64 `json:"amount"`
			Items  float64 `json:"items"`
		}{
			Amount: 799.50,
			Items:  1,
		},
	},
	{
		Timestamp: time.Now().Add(-8 * time.Minute),
		Level:     "INFO",
		TraceID:   "trace_1003",
		UserID:    "user_77",
		OrderID:   "order_5003",
		Message:   "Order created",
		Meta: struct {
			Amount float64 `json:"amount"`
			Items  float64 `json:"items"`
		}{
			Amount: 2499.00,
			Items:  4,
		},
	},
	{
		Timestamp: time.Now().Add(-7 * time.Minute),
		Level:     "INFO",
		TraceID:   "trace_1003",
		UserID:    "user_77",
		OrderID:   "order_5003",
		Message:   "Order confirmed",
		Meta: struct {
			Amount float64 `json:"amount"`
			Items  float64 `json:"items"`
		}{
			Amount: 2499.00,
			Items:  4,
		},
	},
	{
		Timestamp: time.Now().Add(-6 * time.Minute),
		Level:     "INFO",
		TraceID:   "trace_1004",
		UserID:    "user_12",
		OrderID:   "order_5004",
		Message:   "Order shipped",
		Meta: struct {
			Amount float64 `json:"amount"`
			Items  float64 `json:"items"`
		}{
			Amount: 599.99,
			Items:  1,
		},
	},
	{
		Timestamp: time.Now().Add(-5 * time.Minute),
		Level:     "WARN",
		TraceID:   "trace_1005",
		UserID:    "user_88",
		OrderID:   "order_5005",
		Message:   "Delayed order processing",
		Meta: struct {
			Amount float64 `json:"amount"`
			Items  float64 `json:"items"`
		}{
			Amount: 1299.75,
			Items:  3,
		},
	},
	{
		Timestamp: time.Now().Add(-4 * time.Minute),
		Level:     "INFO",
		TraceID:   "trace_1005",
		UserID:    "user_88",
		OrderID:   "order_5005",
		Message:   "Order processed successfully",
		Meta: struct {
			Amount float64 `json:"amount"`
			Items  float64 `json:"items"`
		}{
			Amount: 1299.75,
			Items:  3,
		},
	},
	{
		Timestamp: time.Now().Add(-3 * time.Minute),
		Level:     "ERROR",
		TraceID:   "trace_1006",
		UserID:    "user_33",
		OrderID:   "order_5006",
		Message:   "Order failed due to payment error",
		Meta: struct {
			Amount float64 `json:"amount"`
			Items  float64 `json:"items"`
		}{
			Amount: 1999.99,
			Items:  2,
		},
	},
}
