package pg_storage

import (
	"github.com/ilhamrobyana/online-store-evermos-task/entity"
	"github.com/jinzhu/gorm"
)

type Product struct {
	Client *gorm.DB
}

func (p *Product) Create(product entity.Product) (entity.Product, error) {
	e := p.Client.Create(&product).Error
	return product, e
}
func (p *Product) GetAll() (productList []entity.Product, e error) {
	e = p.Client.
		Find(&productList).
		Order("id ASC").
		Error
	return
}
func (p *Product) GetByID(id uint64) (product entity.Product, e error) {
	e = p.Client.
		Where("id=? ", id).
		Find(&product).
		Error
	return
}
func (p *Product) Update(id uint64, product entity.Product) (entity.Product, error) {
	e := p.Client.
		Model(&product).
		Where("id=?", id).
		Updates(product).
		Where("id=?", id).
		Find(&product).Error
	return product, e
}

func (p *Product) Delete(id uint64) (e error) {
	e = p.Client.
		Where("id=?", id).
		Find(&entity.Product{}).
		Delete(&entity.Product{}).
		Error
	return
}
