package pg_storage

import (
	"fmt"

	"github.com/ilhamrobyana/online-store-evermos-task/entity"
	"github.com/jinzhu/gorm"
)

type Order struct{}

func (o *Order) Create(items []entity.Item, userID uint64) (entity.Order, error) {
	client, e := GetPGClient()
	defer client.Close()

	if e != nil {
		return entity.Order{}, e
	}

	order := entity.Order{UserID: userID}

	tx := client.Begin()
	if checkRollback(tx.Create(&order).Error, tx) {
		return entity.Order{}, e
	}

	var totalItems uint
	for _, v := range items {
		fmt.Println(v.Quantity)
		if checkRollback(tx.Model(&entity.Product{}).Where("id=? AND inventory >= ?", v.ProductID, v.Quantity).Update("inventory", gorm.Expr("inventory - ?", v.Quantity)).Error, tx) {
			return entity.Order{}, e
		}
		if checkRollback(tx.Create(&entity.OrderProduct{OrderID: order.ID, ProductID: v.ProductID, Quantity: v.Quantity}).Error, tx) {
			return entity.Order{}, e
		}
		totalItems += v.Quantity
	}
	order.Items = totalItems
	if checkRollback(tx.Model(&order).Where("id=?", order.ID).Updates(order).Error, tx) {
		return entity.Order{}, e
	}
	if checkRollback(tx.Commit().Error, tx) {
		return entity.Order{}, e
	}
	return order, e
}

func (o *Order) GetAll(userID uint64) (orderList []entity.Order, e error) {
	client, e := GetPGClient()
	defer client.Close()

	if e != nil {
		return
	}

	e = client.
		Where("user_id=?", userID).
		Find(&orderList).
		Order("id ASC").
		Error
	return
}
func (o *Order) GetByID(id uint64) (order entity.Order, e error) {
	client, e := GetPGClient()
	defer client.Close()

	if e != nil {
		return
	}

	e = client.
		Where("id=? ", id).
		Find(&order).
		Error
	return
}
func (o *Order) Update(id uint64, order entity.Order) (entity.Order, error) {
	client, e := GetPGClient()
	defer client.Close()

	if e != nil {
		return order, e
	}

	e = client.
		Model(&order).
		Where("id=?", id).
		Updates(order).
		Where("id=?", id).
		Find(&order).Error
	return order, e
}

func (o *Order) Delete(id uint64) (e error) {
	client, e := GetPGClient()
	defer client.Close()

	if e != nil {
		return
	}
	e = client.
		Where("id=?", id).
		Find(&entity.Order{}).
		Delete(&entity.Order{}).
		Error
	return
}

func checkRollback(e error, tx *gorm.DB) bool {
	if e != nil {
		tx.Rollback()
		return true
	}
	return false
}
