package storage

import (
	"errors"

	"github.com/ilhamrobyana/online-store-evermos-task/entity"
	pg "github.com/ilhamrobyana/online-store-evermos-task/pg_storage"
)

type OrderStorage interface {
	Create(item []entity.Item, userID uint64) (entity.Order, error)
	GetAll(userID uint64) ([]entity.Order, error)
	GetByID(id uint64) (entity.Order, error)
	Update(id uint64, order entity.Order) (entity.Order, error)
	Delete(id uint64) error
}

func GetOrderStorage(n int) (OrderStorage, error) {
	switch n {
	case Postgre:
		return new(pg.Order), nil
	default:
		return nil, errors.New("not implemented yet")
	}
}
