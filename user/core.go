package user

import (
	"errors"

	"github.com/ilhamrobyana/online-store-evermos-task/entity"
	"github.com/ilhamrobyana/online-store-evermos-task/helper"
	"github.com/ilhamrobyana/online-store-evermos-task/storage"
	"golang.org/x/crypto/bcrypt"
)

type core struct {
	userStore storage.UserStorage
}

func (c *core) signup(user entity.User) (response entity.LoginResponse, e error) {
	hashedPassword, e := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if e != nil {
		return
	}
	user.Password = string(hashedPassword)
	createdUser, e := c.userStore.Create(user)

	response.Token, e = helper.GenerateToken(createdUser)
	return
}

func (c *core) login(username, password string) (response entity.LoginResponse, e error) {
	user, e := c.userStore.GetByUsername(username)
	if e != nil {
		return
	}
	e = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if e != nil {
		e = errors.New("Username or Password is wrong")
	}
	response.Token, e = helper.GenerateToken(user)
	return
}
