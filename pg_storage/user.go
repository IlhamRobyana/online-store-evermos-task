package pg_storage

import "github.com/ilhamrobyana/online-store-evermos-task/entity"

type User struct{}

func (u *User) Create(user entity.User) (entity.User, error) {
	client, e := GetPGClient()
	defer client.Close()

	if e != nil {
		return entity.User{}, e
	}
	e = client.Create(&user).Error
	return user, e
}

func (p *User) GetByUsername(username string) (user entity.User, e error) {
	client, e := GetPGClient()
	defer client.Close()
	if e != nil {
		return
	}
	e = client.Where("username = ?", username).First(&user).Error
	return
}
