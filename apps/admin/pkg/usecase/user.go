package usecase

import (
	"context"
	"fmt"
	"time"

	db "github.com/blanc42/becho/pkg/db/sqlc"
	"github.com/blanc42/becho/pkg/domain"
	"github.com/blanc42/becho/pkg/handlers/request"
	"github.com/blanc42/becho/pkg/handlers/response"
	"github.com/blanc42/becho/pkg/utils"
	"github.com/jackc/pgx/v5/pgtype"
)

type UserUsecase interface {
	CreateUser(ctx context.Context, user_request request.CreateUserRequest) (string, error)
	GetAdminUser(ctx context.Context, user_id string) (response.AdminLoginResponse, error)
	GetCustomer(ctx context.Context, id string) (response.CustomerLoginResponse, error)
	UpdateUser(ctx context.Context, id string, user_request request.UpdateUserRequest) (db.User, error)
	DeleteUser(ctx context.Context, id string) error
	ListUsers(ctx context.Context, limit, offset int32) ([]db.User, error)
	AuthenticateUser(ctx context.Context, email, password string) (response.AdminLoginResponse, error)
	CreateAdminUser(ctx context.Context, user_request request.CreateAdminRequest) (string, error)
}

type userUseCase struct {
	userRepo  domain.UserRepository
	storeRepo domain.StoreRepository
}

func NewUserUseCase(userRepo domain.UserRepository, storeRepo domain.StoreRepository) UserUsecase {
	return &userUseCase{userRepo: userRepo, storeRepo: storeRepo}
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

func (u *userUseCase) GetAdminUser(ctx context.Context, user_id string) (response.AdminLoginResponse, error) {
	user, err := u.userRepo.GetUser(ctx, user_id)
	if err != nil {
		return response.AdminLoginResponse{}, fmt.Errorf("failed to get user: %w", err)
	}
	stores, err := u.storeRepo.ListStores(ctx, user_id)

	fmt.Println("I am from the GET ADMIN USER use case <<<============================")
	fmt.Printf("the user is %s with stores %v", user_id, stores)
	if err != nil {
		return response.AdminLoginResponse{}, fmt.Errorf("failed to get stores: %w", err)
	}
	if len(stores) == 0 {
		stores = []db.Store{}
	}

	login_response := response.AdminLoginResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Role:     string(user.Role),
		Stores:   stores,
	}

	return login_response, nil
}

func (u *userUseCase) GetCustomer(ctx context.Context, id string) (response.CustomerLoginResponse, error) {
	user, err := u.userRepo.GetUser(ctx, id)
	if err != nil {
		return response.CustomerLoginResponse{}, fmt.Errorf("failed to get user: %w", err)
	}

	login_response := response.CustomerLoginResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Role:     string(user.Role),
	}

	return login_response, nil
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
	store_id := ctx.Value("store_id")
	users, err := u.userRepo.ListUsers(ctx, db.ListUsersForStoreParams{
		StoreID: pgtype.Text{String: store_id.(string), Valid: true},
		Limit:   limit,
		Offset:  offset,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to list users: %w", err)
	}
	return users, nil
}

func (u *userUseCase) AuthenticateUser(ctx context.Context, email, password string) (response.AdminLoginResponse, error) {
	user, err := u.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return response.AdminLoginResponse{}, fmt.Errorf("failed to get user by email: %w", err)
	}

	stores, err := u.storeRepo.ListStores(ctx, user.ID)

	if len(stores) == 0 {
		stores = []db.Store{}
	}

	login_response := response.AdminLoginResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Role:     string(user.Role),
		Stores:   stores,
	}

	if err != nil {
		return response.AdminLoginResponse{}, fmt.Errorf("failed to get stores: %w", err)
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return response.AdminLoginResponse{}, fmt.Errorf("invalid credentials")
	}

	return login_response, nil
}
