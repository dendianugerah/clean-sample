package model

import "errors"

type User struct {
	ID       int
	Username string
}

func NewUser(username string) (*User, error) {
	if username == "" {
		return nil, errors.New("please enter the username")
	}

	user := &User{
		Username: username,
	}

	return user, nil
}

func (u *User) Set(username string) error {
	if username == "" {
		return errors.New("please enter the username")
	}

	u.Username = username
	return nil
}
