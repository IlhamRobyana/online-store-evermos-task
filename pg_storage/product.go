package pg_storage

import "github.com/ilhamrobyana/online-store-evermos-task/entity"

type Product struct{}

func (p *Product) Create(product entity.Product) (entity.Product, error) {
	client, e := GetPGClient()
	defer client.Close()

	if e != nil {
		return product, e
	}

	e = client.Create(&product).Error
	return product, e
}
func (p *Product) GetAll() (productList []entity.Product, e error) {
	client, e := GetPGClient()
	defer client.Close()

	if e != nil {
		return
	}

	e = client.
		Find(&productList).
		Order("id ASC").
		Error
	return
}
func (p *Product) GetByID(id uint64) (product entity.Product, e error) {
	client, e := GetPGClient()
	defer client.Close()

	if e != nil {
		return
	}

	e = client.
		Where("id=? ", id).
		Find(&product).
		Error
	return
}
func (p *Product) Update(id uint64, product entity.Product) (entity.Product, error) {
	client, e := GetPGClient()
	defer client.Close()

	if e != nil {
		return product, e
	}

	e = client.
		Model(&product).
		Where("id=?", id).
		Updates(product).
		Where("id=?", id).
		Find(&product).Error
	return product, e
}

func (p *Product) Delete(id uint64) (e error) {
	client, e := GetPGClient()
	defer client.Close()

	if e != nil {
		return
	}
	e = client.
		Where("id=?", id).
		Find(&entity.Product{}).
		Delete(&entity.Product{}).
		Error
	return
}
