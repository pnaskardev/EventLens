package utils

import (
	"time"

	"github.com/pnaskardev/EventLens/auth-service/pkg/entities"
)

var AuthLogs = []entities.AuthLog{
	{
		Timestamp: time.Now().Add(-10 * time.Minute),
		Level:     "INFO",
		TraceID:   "trace_001",
		OrderID:   "order_1001",
		Message:   "Payment successful",
		Meta: struct {
			Provider string `json:"provider"`
			Reason   string `json:"reason"`
		}{
			Provider: "stripe",
			Reason:   "",
		},
	},
	{
		Timestamp: time.Now().Add(-9 * time.Minute),
		Level:     "ERROR",
		TraceID:   "trace_002",
		OrderID:   "order_1002",
		Message:   "Payment failed",
		Meta: struct {
			Provider string `json:"provider"`
			Reason   string `json:"reason"`
		}{
			Provider: "razorpay",
			Reason:   "insufficient_funds",
		},
	},
	{
		Timestamp: time.Now().Add(-8 * time.Minute),
		Level:     "WARN",
		TraceID:   "trace_003",
		OrderID:   "order_1003",
		Message:   "Payment retry initiated",
		Meta: struct {
			Provider string `json:"provider"`
			Reason   string `json:"reason"`
		}{
			Provider: "stripe",
			Reason:   "network_timeout",
		},
	},
	{
		Timestamp: time.Now().Add(-7 * time.Minute),
		Level:     "INFO",
		TraceID:   "trace_003",
		OrderID:   "order_1003",
		Message:   "Payment successful after retry",
		Meta: struct {
			Provider string `json:"provider"`
			Reason   string `json:"reason"`
		}{
			Provider: "stripe",
			Reason:   "",
		},
	},
	{
		Timestamp: time.Now().Add(-6 * time.Minute),
		Level:     "ERROR",
		TraceID:   "trace_004",
		OrderID:   "order_1004",
		Message:   "Payment failed",
		Meta: struct {
			Provider string `json:"provider"`
			Reason   string `json:"reason"`
		}{
			Provider: "paypal",
			Reason:   "card_declined",
		},
	},
	{
		Timestamp: time.Now().Add(-5 * time.Minute),
		Level:     "INFO",
		TraceID:   "trace_005",
		OrderID:   "order_1005",
		Message:   "Payment initiated",
		Meta: struct {
			Provider string `json:"provider"`
			Reason   string `json:"reason"`
		}{
			Provider: "razorpay",
			Reason:   "",
		},
	},
	{
		Timestamp: time.Now().Add(-4 * time.Minute),
		Level:     "ERROR",
		TraceID:   "trace_005",
		OrderID:   "order_1005",
		Message:   "Payment failed",
		Meta: struct {
			Provider string `json:"provider"`
			Reason   string `json:"reason"`
		}{
			Provider: "razorpay",
			Reason:   "gateway_error",
		},
	},
	{
		Timestamp: time.Now().Add(-3 * time.Minute),
		Level:     "WARN",
		TraceID:   "trace_006",
		OrderID:   "order_1006",
		Message:   "Slow payment processing",
		Meta: struct {
			Provider string `json:"provider"`
			Reason   string `json:"reason"`
		}{
			Provider: "stripe",
			Reason:   "high_latency",
		},
	},
	{
		Timestamp: time.Now().Add(-2 * time.Minute),
		Level:     "INFO",
		TraceID:   "trace_006",
		OrderID:   "order_1006",
		Message:   "Payment completed",
		Meta: struct {
			Provider string `json:"provider"`
			Reason   string `json:"reason"`
		}{
			Provider: "stripe",
			Reason:   "",
		},
	},
	{
		Timestamp: time.Now().Add(-1 * time.Minute),
		Level:     "ERROR",
		TraceID:   "trace_007",
		OrderID:   "order_1007",
		Message:   "Payment failed",
		Meta: struct {
			Provider string `json:"provider"`
			Reason   string `json:"reason"`
		}{
			Provider: "paypal",
			Reason:   "fraud_suspected",
		},
	},
}
