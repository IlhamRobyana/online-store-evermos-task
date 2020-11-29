package order

import (
	"github.com/ilhamrobyana/online-store-evermos-task/entity"
	pg "github.com/ilhamrobyana/online-store-evermos-task/pg_storage"
)

type core struct {
	orderStorage pg.Order
}

func (c *core) create(items []entity.Item, userID uint64) (createdOrder entity.Order, e error) {
	client, e := pg.GetPGClient()
	defer client.Close()
	if e != nil {
		return
	}
	c.orderStorage.Client = client

	createdOrder, e = c.orderStorage.Create(items, userID)
	return
}

func (c *core) getAll(userID uint64) (orderList []entity.Order, e error) {
	client, e := pg.GetPGClient()
	defer client.Close()
	if e != nil {
		return
	}
	c.orderStorage.Client = client

	orderList, e = c.orderStorage.GetAll(userID)
	return
}

func (c *core) getByID(id uint64) (order entity.Order, e error) {
	client, e := pg.GetPGClient()
	defer client.Close()
	if e != nil {
		return
	}
	c.orderStorage.Client = client

	order, e = c.orderStorage.GetByID(id)
	return
}

func (c *core) update(id uint64, order entity.Order) (entity.Order, error) {
	client, e := pg.GetPGClient()
	defer client.Close()
	if e != nil {
		return entity.Order{}, e
	}
	c.orderStorage.Client = client

	order, e = c.orderStorage.Update(id, order)
	return order, e
}

func (c *core) delete(id uint64) (e error) {
	client, e := pg.GetPGClient()
	defer client.Close()
	if e != nil {
		return
	}
	c.orderStorage.Client = client

	e = c.orderStorage.Delete(id)
	return
}
