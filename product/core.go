package product

import (
	"github.com/ilhamrobyana/online-store-evermos-task/entity"
	pg "github.com/ilhamrobyana/online-store-evermos-task/pg_storage"
)

type core struct {
	productStorage pg.Product
}

func (c *core) create(product entity.Product) (createdProduct entity.Product, e error) {
	client, e := pg.GetPGClient()
	defer client.Close()
	if e != nil {
		return
	}
	c.productStorage.Client = client
	createdProduct, e = c.productStorage.Create(product)
	return
}

func (c *core) getAll() (productList []entity.Product, e error) {
	client, e := pg.GetPGClient()
	defer client.Close()
	if e != nil {
		return
	}
	c.productStorage.Client = client
	productList, e = c.productStorage.GetAll()
	return
}

func (c *core) getByID(id uint64) (product entity.Product, e error) {
	client, e := pg.GetPGClient()
	defer client.Close()
	if e != nil {
		return
	}
	c.productStorage.Client = client
	product, e = c.productStorage.GetByID(id)
	return
}

func (c *core) update(id uint64, product entity.Product) (entity.Product, error) {
	client, e := pg.GetPGClient()
	defer client.Close()
	if e != nil {
		return entity.Product{}, e
	}
	c.productStorage.Client = client
	product, e = c.productStorage.Update(id, product)
	return product, e
}

func (c *core) delete(id uint64) (e error) {
	client, e := pg.GetPGClient()
	defer client.Close()
	if e != nil {
		return
	}
	c.productStorage.Client = client
	e = c.productStorage.Delete(id)
	return
}
