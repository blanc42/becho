package usecase

import (
	"context"
	"fmt"
	"time"

	db "github.com/blanc42/becho/pkg/db/sqlc"
	"github.com/blanc42/becho/pkg/domain"
	"github.com/blanc42/becho/pkg/handlers/request"
	"github.com/blanc42/becho/pkg/utils"
	"github.com/jackc/pgx/v5/pgtype"
)

type UserUsecase interface {
	CreateUser(ctx context.Context, user_request request.CreateUserRequest) (string, error)
	GetUser(ctx context.Context, id string) (db.User, error)
	UpdateUser(ctx context.Context, id string, user_request request.UpdateUserRequest) (db.User, error)
	DeleteUser(ctx context.Context, id string) error
	ListUsers(ctx context.Context, limit, offset int32) ([]db.User, error)
	AuthenticateUser(ctx context.Context, email, password string) (db.User, error)
	CreateAdminUser(ctx context.Context, user_request request.CreateAdminRequest) (string, error)
}

type userUseCase struct {
	userRepo domain.UserRepository
}

func NewUserUseCase(userRepo domain.UserRepository) UserUsecase {
	return &userUseCase{userRepo: userRepo}
}

func (u *userUseCase) CreateAdminUser(ctx context.Context, user_request request.CreateAdminRequest) (string, error) {
	id, err := utils.GenerateShortID()
	if err != nil {
		return "", fmt.Errorf("failed to generate ID: %w", err)
	}

	hashedPassword, err := utils.HashPassword(user_request.Password)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}

	newUser := db.User{
		ID:        id,
		CreatedAt: pgtype.Timestamp{Time: time.Now(), Valid: true},
		UpdatedAt: pgtype.Timestamp{Time: time.Now(), Valid: true},
		Username:  user_request.Username,
		Email:     user_request.Email,
		Password:  hashedPassword,
		Role:      "admin",
		StoreID:   pgtype.Text{},
	}

	createdID, err := u.userRepo.CreateUser(ctx, &newUser)
	if err != nil {
		return "", fmt.Errorf("failed to create user: %w", err)
	}

	return createdID, nil
}

func (u *userUseCase) CreateUser(ctx context.Context, user_request request.CreateUserRequest) (string, error) {
	id, err := utils.GenerateShortID()
	if err != nil {
		return "", fmt.Errorf("failed to generate ID: %w", err)
	}

	hashedPassword, err := utils.HashPassword(user_request.Password)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}

	newUser := db.User{
		ID:        id,
		CreatedAt: pgtype.Timestamp{Time: time.Now(), Valid: true},
		UpdatedAt: pgtype.Timestamp{Time: time.Now(), Valid: true},
		Username:  user_request.Username,
		Email:     user_request.Email,
		Password:  hashedPassword,
		Role:      "customer",
		StoreID:   pgtype.Text{String: user_request.StoreID, Valid: true},
	}

	createdID, err := u.userRepo.CreateUser(ctx, &newUser)
	if err != nil {
		return "", fmt.Errorf("failed to create user: %w", err)
	}

	return createdID, nil
}

func (u *userUseCase) GetUser(ctx context.Context, id string) (db.User, error) {
	user, err := u.userRepo.GetUser(ctx, id)
	if err != nil {
		return db.User{}, fmt.Errorf("failed to get user: %w", err)
	}
	return user, nil
}

func (u *userUseCase) UpdateUser(ctx context.Context, id string, user_request request.UpdateUserRequest) (db.User, error) {
	existingUser, err := u.userRepo.GetUser(ctx, id)
	if err != nil {
		return db.User{}, fmt.Errorf("failed to get user: %w", err)
	}

	updateParams := db.UpdateUserParams{
		ID:        id,
		UpdatedAt: pgtype.Timestamp{Time: time.Now(), Valid: true},
	}

	if user_request.Username != nil {
		updateParams.Username = *user_request.Username
	} else {
		updateParams.Username = existingUser.Username
	}

	if user_request.Email != nil {
		updateParams.Email = *user_request.Email
	} else {
		updateParams.Email = existingUser.Email
	}

	if user_request.Password != nil {
		hashedPassword, err := utils.HashPassword(*user_request.Password)
		if err != nil {
			return db.User{}, fmt.Errorf("failed to hash password: %w", err)
		}
		updateParams.Password = hashedPassword
	} else {
		updateParams.Password = existingUser.Password
	}

	updateParams.Role = existingUser.Role
	updateParams.StoreID = existingUser.StoreID

	updatedUser, err := u.userRepo.UpdateUser(ctx, updateParams)
	if err != nil {
		return db.User{}, fmt.Errorf("failed to update user: %w", err)
	}

	return updatedUser, nil
}

func (u *userUseCase) DeleteUser(ctx context.Context, id string) error {
	err := u.userRepo.DeleteUser(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}

func (u *userUseCase) ListUsers(ctx context.Context, limit, offset int32) ([]db.User, error) {
	users, err := u.userRepo.ListUsers(ctx, db.ListUsersParams{Limit: limit, Offset: offset})
	if err != nil {
		return nil, fmt.Errorf("failed to list users: %w", err)
	}
	return users, nil
}

func (u *userUseCase) AuthenticateUser(ctx context.Context, email, password string) (db.User, error) {
	user, err := u.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return db.User{}, fmt.Errorf("failed to get user by email: %w", err)
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return db.User{}, fmt.Errorf("invalid credentials")
	}

	return user, nil
}
