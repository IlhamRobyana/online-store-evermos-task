package user

import (
	"errors"

	"github.com/ilhamrobyana/online-store-evermos-task/entity"
	"github.com/ilhamrobyana/online-store-evermos-task/helper"
	pg "github.com/ilhamrobyana/online-store-evermos-task/pg_storage"

	"golang.org/x/crypto/bcrypt"
)

type core struct {
	userStore pg.User
}

func (c *core) signup(user entity.User) (response entity.LoginResponse, e error) {
	client, e := pg.GetPGClient()
	defer client.Close()

	if e != nil {
		return entity.LoginResponse{}, e
	}

	c.userStore.Client = client

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
	client, e := pg.GetPGClient()
	defer client.Close()

	if e != nil {
		return entity.LoginResponse{}, e
	}

	c.userStore.Client = client
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
