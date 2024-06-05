package domain

import (
	"errors"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	id       uuid.UUID
	name     string
	password string
}

func (u *User) ID() uuid.UUID {
	return u.id
}

func (u *User) Name() string {
	return u.name
}

func (u *User) SetName(value string) error {
	if len(value) < 3 {
		return errors.New("name is too small")
	}
	u.name = value
	return nil
}

func (u *User) Password() string {
	return u.password
}

func (u *User) SetPassword(value string) error {
	if len(value) < 6 {
		return errors.New("password is too simple")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.MinCost)
	if err != nil {
		return err
	}
	u.password = string(hashed)
	return nil
}

func (u *User) ComparePassword(value string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.password), []byte(value)) == nil
}

// NewUser creates a new user model with provided values
func NewUser(name string, password string) (User, error) {
	user := User{id: uuid.New()}

	nameErr := user.SetName(name)
	passwordErr := user.SetPassword(password)
	if err := errors.Join(nameErr, passwordErr); err != nil {
		return User{}, err
	}

	return user, nil
}

// NewRawUser generates a user struct with provided data.
//
// NOTE: This method does not encrypt password.
// NOTE: This method does not validates the input data.
//
// NOTE: This method should not be used in business logic.
// Its purpose is to create model from the storage.
func NewRawUser(id uuid.UUID, name string, password string) User {
	return User{
		id:       id,
		name:     name,
		password: password,
	}
}
