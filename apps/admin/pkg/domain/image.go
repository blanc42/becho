package domain

import (
	"context"

	db "github.com/blanc42/becho/apps/admin/pkg/db/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

type ImageRepository struct {
	db *db.DbStore
}

type ImageRepositoryInterface interface {
	CreateImage(ctx context.Context, arg db.CreateImageParams) (db.Image, error)
}

func NewImageRepository(db *db.DbStore) ImageRepository {
	return ImageRepository{db: db}
}

func (r *ImageRepository) CreateImage(ctx context.Context, arg db.CreateImageParams) (db.Image, error) {
	return r.db.CreateImage(ctx, arg)
}

func (r *ImageRepository) DeleteImage(ctx context.Context, image_id pgtype.Int4) error {
	return r.db.DeleteImage(ctx, image_id)
}
