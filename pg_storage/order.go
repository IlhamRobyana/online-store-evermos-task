package pg_storage

import "github.com/ilhamrobyana/online-store-evermos-task/entity"

type Order struct{}

func (o *Order) Create(order entity.Order) (entity.Order, error) {
	client, e := GetPGClient()
	defer client.Close()

	if e != nil {
		return order, e
	}

	e = client.Create(&order).Error
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
