package domain

import (
	"context"
	"github.com/go-gorp/gorp"
	"github.com/google/uuid"
	"github.com/taroodr/til/go/graphql-model-relation-test/model"
)

type UserID string

type User struct {
	ID UserID
	Name string
}

type UserRepository interface {
	Get(ctx context.Context, id UserID) (*User, error)
	Create(ctx context.Context, user *User) (*User, error)
}

func NewUserRepository(db *gorp.DbMap) UserRepository {
	return &userRepository{
		db: db,
	}
}

var _ UserRepository = (*userRepository)(nil)

type userRepository struct {
	db *gorp.DbMap
}

func (repo *userRepository) Get(ctx context.Context, id UserID) (*User, error) {
	var user model.User
	err := repo.db.SelectOne(&user, "select * from user where id=?", id)
	if err != nil {
		return nil, err
	}
	return &User{
		ID: UserID(user.ID),
		Name: user.Name,
	}, nil
}


func (repo *userRepository) Create(ctx context.Context, user *User) (*User, error) {
	if user.ID != "" {
		return nil, ErrBadRequest
	}

	id := uuid.New().String()

	um := &model.User{
		ID: id,
		Name: user.Name,
	}
	err := repo.db.Insert(um)

	if err != nil {
		return nil, err
	}
	return &User{
		ID: UserID(um.ID),
		Name: um.Name,
	}, nil
}