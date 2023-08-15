package store

import (
	"crypto/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// we gonna store hashed password with random salt.
// Here
// We added pg:"-" for password validation. what that means? this is simply defines that there is no password column in "users" table.This field is still required because user needs to provide it, but it will not be saved in database
// Instead of that, we will generate random Salt and use Password field to generate HashedPassword. Then we will store Salt and HashedPassword in database.
// When authenticating user with Username and Password, we can use Salt from database to calculate HashedPassword based on provided Password and then compare it with HashedPassword stored in database.
// We also added json:"-" for HashedPassword, and Salt fields so they will not be sent to fronted in JSON response and will also be ignored if sent to backend as JSON.
// User struct is now ready to go, but we also need to update our migration file:
type User struct {
	ID             int
	Username       string `binding:"required,min=5,max=30"`
	Password       string `pg:"-" binding:"required,min=3,max=32"`
	HashedPassword []byte `json:"-"`
	Salt           []byte `json:"-"`
	CreatedAt      time.Time
	ModifiedAt     time.Time
}

func AddUser(user *User) error {
	salt, err := GenerateSalt()
	if err != nil {
		return err
	}
	toHash := append([]byte(user.Password), salt...)
	hashedPassword, err := bcrypt.GenerateFromPassword(toHash, bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Salt = salt
	user.HashedPassword = hashedPassword

	_, err = db.Model(user).Returning("*").Insert()

	if err != nil {
		return err
	}

	return err
}

func Authenticate(username, password string) (*User, error) {
	user := new(User)

	if err := db.Model(user).Where(
		"username = ?", username).Select(); err != nil {
		return nil, err
	}

	salted := append([]byte(password), user.Salt...)
	if err := bcrypt.CompareHashAndPassword(user.HashedPassword, salted); err != nil {
		return nil, err
	}

	return user, nil
}

func GenerateSalt() ([]byte, error) {
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return nil, err
	}
	return salt, nil
}
