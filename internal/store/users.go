package store

import "errors"

type User struct {
	Username string `binding:"required,min=5,max=30"`
	Password string `binding:"required,min=3,max=32"`
}

func AddUser(user *User) error {
	_, err := db.Model(user).Returning("*").Insert()

	if err != nil {
		return err
	}

	return nil
}

func Authenticate(username, password string) (*User, error) {
	user := new(User)

	if err := db.Model(user).Where("username = ?", username).Select(); err != nil {
		return nil, err
	}

	if password != user.Password {
		return nil, errors.New("password not valid")
	}

	return user, nil
}
