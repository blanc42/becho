// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: images.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createImage = `-- name: CreateImage :one
INSERT INTO images (created_at, updated_at, title, image_id)
VALUES ($1, $2, $3, $4)
RETURNING created_at, updated_at, title, image_id, id
`

type CreateImageParams struct {
	CreatedAt pgtype.Timestamp `json:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
	Title     pgtype.Text      `json:"title"`
	ImageID   string           `json:"image_id"`
}

func (q *Queries) CreateImage(ctx context.Context, arg CreateImageParams) (Image, error) {
	row := q.db.QueryRow(ctx, createImage,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Title,
		arg.ImageID,
	)
	var i Image
	err := row.Scan(
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Title,
		&i.ImageID,
		&i.ID,
	)
	return i, err
}

const deleteImage = `-- name: DeleteImage :exec
DELETE FROM images
WHERE id = $1
`

func (q *Queries) DeleteImage(ctx context.Context, id pgtype.Int4) error {
	_, err := q.db.Exec(ctx, deleteImage, id)
	return err
}

const updateImage = `-- name: UpdateImage :one
UPDATE images
SET title = $2, image_id = $3, updated_at = $4
WHERE id = $1
RETURNING created_at, updated_at, title, image_id, id
`

type UpdateImageParams struct {
	ID        pgtype.Int4      `json:"id"`
	Title     pgtype.Text      `json:"title"`
	ImageID   string           `json:"image_id"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
}

func (q *Queries) UpdateImage(ctx context.Context, arg UpdateImageParams) (Image, error) {
	row := q.db.QueryRow(ctx, updateImage,
		arg.ID,
		arg.Title,
		arg.ImageID,
		arg.UpdatedAt,
	)
	var i Image
	err := row.Scan(
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Title,
		&i.ImageID,
		&i.ID,
	)
	return i, err
}
