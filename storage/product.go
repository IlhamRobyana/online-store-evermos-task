package storage

import (
	"errors"

	"github.com/ilhamrobyana/online-store-evermos-task/entity"
	pg "github.com/ilhamrobyana/online-store-evermos-task/pg_storage"
)

type ProductStorage interface {
	Create(product entity.Product) (entity.Product, error)
	GetAll() ([]entity.Product, error)
	GetByID(id uint64) (entity.Product, error)
	Update(id uint64, product entity.Product) (entity.Product, error)
	Delete(id uint64) error
}

func GetProductStorage(n int) (ProductStorage, error) {
	switch n {
	case Postgre:
		return new(pg.Product), nil
	default:
		return nil, errors.New("not implemented yet")
	}
}
