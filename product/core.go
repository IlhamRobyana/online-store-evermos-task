package product

import (
	"github.com/ilhamrobyana/online-store-evermos-task/entity"
	"github.com/ilhamrobyana/online-store-evermos-task/storage"
)

type core struct {
	productStorage storage.ProductStorage
}

func (c *core) create(product entity.Product) (createdProduct entity.Product, e error) {
	createdProduct, e = c.productStorage.Create(product)
	return
}

func (c *core) getAll() (productList []entity.Product, e error) {
	productList, e = c.productStorage.GetAll()
	return
}

func (c *core) getByID(id uint64) (product entity.Product, e error) {
	product, e = c.productStorage.GetByID(id)
	return
}

func (c *core) update(id uint64, product entity.Product) (entity.Product, error) {
	product, e := c.productStorage.Update(id, product)
	return product, e
}

func (c *core) delete(id uint64) (e error) {
	e = c.productStorage.Delete(id)
	return
}
