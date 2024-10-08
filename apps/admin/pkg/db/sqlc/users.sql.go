// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: users.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, username, email, password, role, store_id, image_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING id, created_at, updated_at, username, email, password, role, store_id, image_id
`

type CreateUserParams struct {
	ID        string           `json:"id"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
	Username  string           `json:"username"`
	Email     string           `json:"email"`
	Password  string           `json:"password"`
	Role      UserRole         `json:"role"`
	StoreID   pgtype.Text      `json:"store_id"`
	ImageID   pgtype.Int4      `json:"image_id"`
}

// Users
func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Username,
		arg.Email,
		arg.Password,
		arg.Role,
		arg.StoreID,
		arg.ImageID,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.Role,
		&i.StoreID,
		&i.ImageID,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id string) error {
	_, err := q.db.Exec(ctx, deleteUser, id)
	return err
}

const getUser = `-- name: GetUser :one
SELECT id, created_at, updated_at, username, email, password, role, store_id, image_id FROM users
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, id string) (User, error) {
	row := q.db.QueryRow(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.Role,
		&i.StoreID,
		&i.ImageID,
	)
	return i, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, created_at, updated_at, username, email, password, role, store_id, image_id FROM users
WHERE email = $1 LIMIT 1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRow(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.Role,
		&i.StoreID,
		&i.ImageID,
	)
	return i, err
}

const getUserByUsername = `-- name: GetUserByUsername :one
SELECT id, created_at, updated_at, username, email, password, role, store_id, image_id FROM users
WHERE username = $1 LIMIT 1
`

func (q *Queries) GetUserByUsername(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRow(ctx, getUserByUsername, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.Role,
		&i.StoreID,
		&i.ImageID,
	)
	return i, err
}

const listUsersForStore = `-- name: ListUsersForStore :many
SELECT id, created_at, updated_at, username, email, password, role, store_id, image_id FROM users
WHERE store_id = $1
ORDER BY created_at
LIMIT $2 OFFSET $3
`

type ListUsersForStoreParams struct {
	StoreID pgtype.Text `json:"store_id"`
	Limit   int32       `json:"limit"`
	Offset  int32       `json:"offset"`
}

func (q *Queries) ListUsersForStore(ctx context.Context, arg ListUsersForStoreParams) ([]User, error) {
	rows, err := q.db.Query(ctx, listUsersForStore, arg.StoreID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Username,
			&i.Email,
			&i.Password,
			&i.Role,
			&i.StoreID,
			&i.ImageID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUser = `-- name: UpdateUser :one
UPDATE users
SET username = $2, email = $3, password = $4, updated_at = $5, image_id = $6
WHERE id = $1
RETURNING id, created_at, updated_at, username, email, password, role, store_id, image_id
`

type UpdateUserParams struct {
	ID        string           `json:"id"`
	Username  string           `json:"username"`
	Email     string           `json:"email"`
	Password  string           `json:"password"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
	ImageID   pgtype.Int4      `json:"image_id"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, updateUser,
		arg.ID,
		arg.Username,
		arg.Email,
		arg.Password,
		arg.UpdatedAt,
		arg.ImageID,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.Role,
		&i.StoreID,
		&i.ImageID,
	)
	return i, err
}
