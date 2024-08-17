package domain

import (
	"context"
	"fmt"

	db "github.com/blanc42/becho/pkg/db/sqlc"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *db.User) (string, error)
	GetUser(ctx context.Context, id string) (db.User, error)
	UpdateUser(ctx context.Context, arg db.UpdateUserParams) (db.User, error)
	DeleteUser(ctx context.Context, id string) error
	ListUsers(ctx context.Context, arg db.ListUsersForStoreParams) ([]db.User, error)
	GetUserByEmail(ctx context.Context, email string) (db.User, error)
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
		return "", err
	}
	fmt.Println("created user: ", createdUser.ID)
	return createdUser.ID, nil
}

func (ur *userRepository) GetUser(ctx context.Context, id string) (db.User, error) {
	return ur.db.GetUser(ctx, id)
}

func (ur *userRepository) UpdateUser(ctx context.Context, arg db.UpdateUserParams) (db.User, error) {
	return ur.db.UpdateUser(ctx, arg)
}

func (ur *userRepository) DeleteUser(ctx context.Context, id string) error {
	return ur.db.DeleteUser(ctx, id)
}

func (ur *userRepository) ListUsers(ctx context.Context, arg db.ListUsersForStoreParams) ([]db.User, error) {
	return ur.db.ListUsersForStore(ctx, arg)
}

func (ur *userRepository) GetUserByEmail(ctx context.Context, email string) (db.User, error) {
	return ur.db.GetUserByEmail(ctx, email)
}
