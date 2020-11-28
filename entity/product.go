package entity

type Product struct {
	ID        uint64 `json:"id" gorm:"primary_key;type:bigserial"`
	Name      string `json:"name"`
	Inventory int    `json:"inventory"`
	Timestamp
}
