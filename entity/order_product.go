package entity

// OrderProduct is an entity used to store relation of an order and a product
type OrderProduct struct {
	ID        uint64   `json:"id" gorm:"primary_key;type:bigserial"`
	Order     *Order   `json:"order"`
	OrderID   uint64   `json:"order_id" gorm:"foreign_key;type:bigserial"`
	Product   *Product `json:"product"`
	ProductID uint64   `json:"product_id" gorm:"foreign_key;type:bigserial"`
	Quantity  uint     `json:"quantity"`
}
