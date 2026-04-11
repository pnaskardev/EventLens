package entities

import "time"

// {
//   "timestamp": "2026-04-11T18:32:10Z",
//   "service": "order-service",
//   "level": "INFO",
//   "trace_id": "trace_xyz789",
//   "user_id": "user_42",
//   "order_id": "order_567",
//   "message": "Order created",
//   "meta": {
//     "amount": 1200,
//     "items": 3
//   }
// }

type PaymentLog struct {
	Timestamp time.Time `json:"timestamp"`
	Level     string    `json:"level"`
	TraceID   string    `json:"trace_id"`
	OrderID   string    `json:"order_id"`
	Message   string    `json:"Message"`
	Meta      struct {
		Amount float64 `json:"amount"`
		Items  int64   `json:"items"`
	}
}
