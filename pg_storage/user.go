package pg_storage

import (
	"github.com/ilhamrobyana/online-store-evermos-task/entity"
	"github.com/jinzhu/gorm"
)

type User struct {
	Client *gorm.DB
}

func (u *User) Create(user entity.User) (entity.User, error) {
	e := u.Client.Create(&user).Error
	return user, e
}

func (u *User) GetByUsername(username string) (user entity.User, e error) {
	e = u.Client.Where("username = ?", username).First(&user).Error
	return
}
