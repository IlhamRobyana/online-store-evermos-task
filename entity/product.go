package entity

import "time"

type Product struct {
	ID            uint64    `json:"id" gorm:"primary_key;type:bigserial"`
	Name          string    `json:"name"`
	Inventory     uint      `json:"inventory"`
	Price         float64   `json:"price"`
	Discount      float64   `json:"discount"`
	DiscountUntil time.Time `json:"discount_until"`
}
