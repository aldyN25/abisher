package param

import "time"

type Request struct {
	RequestID int64    `json:"request_id" validate:"required"`
	Data      []Record `json:"data" validate:"required"`
}

type Record struct {
	ID        int64     `json:"id" validate:"required"`
	Customers string    `json:"customers" validate:"required"`
	Quantity  int       `json:"quantity" validate:"required"`
	Price     float64   `json:"price" validate:"required"`
	Timestamp time.Time `json:"timestamp"`
}
