package verification

import (
	cbConfig "cab/config"
	"errors"

	pg "gopkg.in/pg.v4"
)

//User structure body
type User struct {
	TableName struct{} `sql:"users" json:"-"`
	ID        string   `sql:"id" json:"id"`
	UserName  string   `sql:"user_name" json:"user_name"`
	Password  string   `sql:"password" json:"password"`
	Token     string   `sql:"token" json:"token"`
}

//GetUserCreds returns users
func (user *User) GetUserCreds(c cbConfig.Config, userName string, pass string) error {

	err := c.PG.Model(&User{}).
		Where("user_name = ? and password=?", userName, pass).
		Select(user)
	return err
}

//UserAuthentication verifies user credentials
func (user *User) UserAuthentication(c cbConfig.Config, token string, userName string, password string) error {
	err := user.GetUserCreds(c, userName, password)
	if err == pg.ErrNoRows {
		return errors.New("Invalid Username or Password")

	}
	if token != user.Token {
		return errors.New("Token Error")
	}
	return nil
}
