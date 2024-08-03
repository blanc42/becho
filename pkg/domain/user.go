package domain

import (
	"context"
	"fmt"

	db "github.com/blanc42/becho/pkg/db/sqlc"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *db.User) (string, error)
}

type userRepository struct {
	db *db.DbStore
}

func NewUserRepository(db *db.DbStore) UserRepository {
	return &userRepository{db: db}
}

func (ur *userRepository) CreateUser(ctx context.Context, user *db.User) (string, error) {
	createdUser, err := ur.db.CreateUser(ctx, db.CreateUserParams(*user))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("created user: ", createdUser.ID)
	return createdUser.ID, nil
}
