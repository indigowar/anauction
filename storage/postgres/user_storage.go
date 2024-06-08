package postgres

import (
	"context"
	"net/mail"
	"net/url"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/indigowar/anauction/domain/models"
	"github.com/indigowar/anauction/domain/service"
	"github.com/indigowar/anauction/storage/postgres/data"
)

type UserStorage struct {
	queries *data.Queries
}

var _ service.UserStorage = &UserStorage{}

// Add implements service.UserStorage.
func (u *UserStorage) Add(ctx context.Context, user models.User) error {
	err := u.queries.InsertUser(ctx, data.InsertUserParams{
		ID:       pgtype.UUID{Bytes: user.ID(), Valid: true},
		Name:     user.Name(),
		Email:    user.Email().String(),
		Password: user.Password(),
	})

	if err != nil {
		// TODO: Handle the error properly
	}

	return nil
}

// Delete implements service.UserStorage.
func (u *UserStorage) Delete(ctx context.Context, id uuid.UUID) error {
	err := u.queries.DeleteUser(ctx, pgtype.UUID{Bytes: id, Valid: true})
	if err != nil {
		// TODO: Handle the error properly
	}
	return nil
}

// GetByEmail implements service.UserStorage.
func (u *UserStorage) GetByEmail(ctx context.Context, email *mail.Address) (models.User, error) {
	object, err := u.queries.GetByEmail(ctx, email.String())
	if err != nil {
		// TODO: Handle the error properly
	}
	return u.toModel(object), nil
}

// GetByID implements service.UserStorage.
func (u *UserStorage) GetByID(ctx context.Context, id uuid.UUID) (models.User, error) {
	object, err := u.queries.GetByID(ctx, pgtype.UUID{Bytes: id, Valid: true})
	if err != nil {
		// TODO: Handle the error properly
	}
	return u.toModel(object), nil
}

// Update implements service.UserStorage.
func (u *UserStorage) Update(ctx context.Context, user models.User) error {
	var image string
	if user.Image() == nil {
		image = ""
	} else {
		image = user.Image().String()
	}

	err := u.queries.UpdateUser(ctx, data.UpdateUserParams{
		ID:       pgtype.UUID{Bytes: user.ID(), Valid: true},
		Name:     user.Name(),
		Email:    user.Email().String(),
		Password: user.Password(),
		Image: pgtype.Text{
			String: image,
			Valid:  user.Image() != nil,
		},
	})

	if err != nil {
		// TODO: handle error properly
	}
	return nil
}

func (u *UserStorage) toModel(obj data.User) models.User {
	email, _ := mail.ParseAddress(obj.Email)
	user := models.NewRawUser(
		obj.ID.Bytes,
		obj.Name,
		email,
		nil,
		obj.Password,
	)
	if obj.Image.String != "" {
		url, _ := url.Parse(obj.Image.String)
		user.SetImage(url)
	}
	return user
}

func NewUserStorage(conn *pgx.Conn) *UserStorage {
	return &UserStorage{
		queries: data.New(conn),
	}
}
