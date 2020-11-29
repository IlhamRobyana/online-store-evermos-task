package order

import (
	"github.com/ilhamrobyana/online-store-evermos-task/entity"
	"github.com/ilhamrobyana/online-store-evermos-task/storage"
)

type core struct {
	orderStorage storage.OrderStorage
}

func (c *core) create(items []entity.Item, userID uint64) (createdOrder entity.Order, e error) {
	createdOrder, e = c.orderStorage.Create(items, userID)
	return
}

func (c *core) getAll(userID uint64) (orderList []entity.Order, e error) {
	orderList, e = c.orderStorage.GetAll(userID)
	return
}

func (c *core) getByID(id uint64) (order entity.Order, e error) {
	order, e = c.orderStorage.GetByID(id)
	return
}

func (c *core) update(id uint64, order entity.Order) (entity.Order, error) {
	order, e := c.orderStorage.Update(id, order)
	return order, e
}

func (c *core) delete(id uint64) (e error) {
	e = c.orderStorage.Delete(id)
	return
}
