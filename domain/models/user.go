package models

import (
	"errors"
	"fmt"
	"net/mail"
	"net/url"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidPassword = errors.New("password is invalid")
	ErrInvalidUsername = errors.New("username is invalid")
)

type User struct {
	id       uuid.UUID
	name     string
	email    *mail.Address
	image    *url.URL
	password string
}

func (u *User) ID() uuid.UUID        { return u.id }
func (u *User) Name() string         { return u.name }
func (u *User) Email() *mail.Address { return u.email }
func (u *User) Image() *url.URL      { return u.image }
func (u *User) Password() string     { return u.password }

func (u *User) SetName(value string) error {
	if len(value) < 4 {
		return &UserValidationError{
			Err:     ErrInvalidUsername,
			Message: "username should be at least 4 characters",
		}
	}
	u.name = value
	return nil
}

func (u *User) SetEmail(value *mail.Address) {
	u.email = value
}

func (u *User) SetImage(value *url.URL) {
	u.image = value
}

func (u *User) SetPassword(value string) error {
	if len(value) < 6 {
		return &UserValidationError{
			Err:     ErrInvalidPassword,
			Message: "user password should be at least 6 characters",
		}
	}
	hashed, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.MinCost)
	if err != nil {
		return err
	}
	u.password = string(hashed)
	return nil
}

func NewUser(
	name string,
	email *mail.Address,
	image *url.URL,
	password string,
) (User, error) {
	user := User{id: uuid.New()}

	nameErr := user.SetName(name)
	passwordErr := user.SetPassword(password)

	user.SetEmail(email)
	user.SetImage(image)

	if nameErr != nil || passwordErr != nil {
		if nameErr == nil {
			return User{}, passwordErr
		}

		if passwordErr == nil {
			return User{}, nameErr
		}

		return User{}, &AggregatedError{
			Errors: []error{nameErr, passwordErr},
		}
	}

	return user, nil
}

func NewRawUser(
	id uuid.UUID,
	name string,
	email *mail.Address,
	image *url.URL,
	password string,
) User {
	return User{
		id:       id,
		name:     name,
		email:    email,
		image:    image,
		password: password,
	}
}

type UserValidationError struct {
	Err     error
	Message string
}

func (uev *UserValidationError) Error() string {
	return fmt.Errorf("%w: %s", uev.Err, uev.Message).Error()
}
