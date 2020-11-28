package entity

type Order struct {
	ID     uint64 `json:"id" gorm:"primary_key;type:bigserial"`
	UserID uint64 `json:"user_id" gorm:"foreign_key;type:bigserial"`
	Items  int    `json:"items"`
	Timestamp
}
