package pg_storage

import (
	"github.com/ilhamrobyana/online-store-evermos-task/entity"
	"github.com/jinzhu/gorm"
)

type Order struct {
	Client *gorm.DB
}

func (o *Order) Create(items []entity.Item, userID uint64) (entity.Order, error) {
	order := entity.Order{UserID: userID}

	tx := o.Client.Begin()
	e := checkRollback(tx.Create(&order).Error, tx)
	if e != nil {
		return entity.Order{}, e
	}
	var totalItems uint
	for _, v := range items {
		product := entity.Product{}
		e = checkRollback(tx.Raw(`UPDATE products
		SET inventory = inventory - ?
		WHERE id=? AND inventory >=?
		RETURNING id`,
			v.Quantity, v.ProductID, v.Quantity).Scan(&product).Error, tx)
		if e != nil {
			return entity.Order{}, e
		}
		e = checkRollback(tx.Create(&entity.OrderProduct{OrderID: order.ID, ProductID: v.ProductID, Quantity: v.Quantity}).Error, tx)
		if e != nil {
			return entity.Order{}, e
		}
		totalItems += v.Quantity
	}
	order.Items = totalItems
	checkRollback(tx.Model(&order).Where("id=?", order.ID).Updates(order).Error, tx)
	if e != nil {
		return entity.Order{}, e
	}
	checkRollback(tx.Commit().Error, tx)
	if e != nil {
		return entity.Order{}, e
	}
	return order, e
}

func (o *Order) GetAll(userID uint64) (orderList []entity.Order, e error) {
	e = o.Client.
		Where("user_id=?", userID).
		Find(&orderList).
		Order("id ASC").
		Error
	return
}
func (o *Order) GetByID(id uint64) (order entity.Order, e error) {
	e = o.Client.
		Where("id=? ", id).
		Find(&order).
		Error
	return
}
func (o *Order) Update(id uint64, order entity.Order) (entity.Order, error) {
	e := o.Client.
		Model(&order).
		Where("id=?", id).
		Updates(order).
		Where("id=?", id).
		Find(&order).Error
	return order, e
}

func (o *Order) Delete(id uint64) (e error) {
	e = o.Client.
		Where("id=?", id).
		Find(&entity.Order{}).
		Delete(&entity.Order{}).
		Error
	return
}

func checkRollback(e error, tx *gorm.DB) error {
	if e != nil {
		tx.Rollback()
	}
	return e
}
