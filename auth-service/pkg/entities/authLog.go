package entities

import "time"

type AuthLog struct {
	Timestamp time.Time `json:"timestamp"`
	Level     string    `json:"level"`
	TraceID   string    `json:"trace_id"`
	OrderID   string    `json:"order_id"`
	Message   string    `json:"message"`
	Meta      struct {
		Provider string `json:"provider"`
		Reason   string `json:"reason"`
	}
}
