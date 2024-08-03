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
	CreateAdminUser(ctx context.Context, user request.CreateUserRequest) (string, error)
}

type userUseCase struct {
	userRepo domain.UserRepository
}

func NewUserUseCase(userRepo domain.UserRepository) UserUsecase {
	return &userUseCase{userRepo: userRepo}
}

func (u *userUseCase) CreateAdminUser(ctx context.Context, user_request request.CreateUserRequest) (string, error) {

	id, err := utils.GenerateShortID()
	if err != nil {
		fmt.Println(err)
	}

	// generate hashed password
	// hashedPassword, err := utils.HashPassword(user_request.Password)
	// if err != nil {
	// 	return err
	// }

	var newUser = db.User{
		ID:        id,
		CreatedAt: pgtype.Timestamp{Time: time.Now(), Valid: true},
		UpdatedAt: pgtype.Timestamp{Time: time.Now(), Valid: true},
		Username:  user_request.Username,
		Email:     user_request.Email,
		Password:  user_request.Password,
		Role:      "admin",
		StoreID:   pgtype.Text{},
	}

	_, err = u.userRepo.CreateUser(ctx, &newUser)
	if err != nil {
		return "", err
	}

	return id, nil
}
