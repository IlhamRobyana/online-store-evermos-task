package entity

type Order struct {
	ID     uint64 `json:"id" gorm:"primary_key;type:bigserial"`
	UserID uint64 `json:"user_id" gorm:"foreign_key;type:bigserial"`
	Items  uint   `json:"items"`
	Timestamp
}

type Item struct {
	ProductID uint64 `json:"product_id" form:"product_id"`
	Quantity  uint   `json:"quantity" form:"quantity"`
}

type OrderCreateRequest struct {
	Items []Item `json:"items" form:"items"`
}
