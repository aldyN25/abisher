package model

import "time"

type Product struct {
	ID        int64     `gorm:"column:id"`
	RequestID int64     `gorm:"column:request_id"`
	Customers string    `gorm:"column:customers"`
	Quantity  int       `gorm:"column:quantity"`
	Price     float64   `gorm:"column:price"`
	CreatedAt time.Time `gorm:"column:created_at"`
}
