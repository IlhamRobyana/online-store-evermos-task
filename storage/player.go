package storage

import (
	"errors"

	"github.com/ilhamrobyana/online-store-evermos-task/entity"
	pg "github.com/ilhamrobyana/online-store-evermos-task/pg_storage"
)

type UserStorage interface {
	Create(user entity.User) (entity.User, error)
	GetByUsername(username string) (entity.User, error)
}

func GetUserStorage(n int) (UserStorage, error) {
	switch n {
	case Postgre:
		return new(pg.User), nil
	default:
		return nil, errors.New("not implemented yet")
	}
}
