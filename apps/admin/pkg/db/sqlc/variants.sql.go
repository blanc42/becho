// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: variants.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createVariant = `-- name: CreateVariant :one
INSERT INTO variants (id, created_at, updated_at, name, label, description, store_id)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING id, created_at, updated_at, name, label, description, store_id
`

type CreateVariantParams struct {
	ID          string           `json:"id"`
	CreatedAt   pgtype.Timestamp `json:"created_at"`
	UpdatedAt   pgtype.Timestamp `json:"updated_at"`
	Name        string           `json:"name"`
	Label       string           `json:"label"`
	Description pgtype.Text      `json:"description"`
	StoreID     string           `json:"store_id"`
}

func (q *Queries) CreateVariant(ctx context.Context, arg CreateVariantParams) (Variant, error) {
	row := q.db.QueryRow(ctx, createVariant,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Name,
		arg.Label,
		arg.Description,
		arg.StoreID,
	)
	var i Variant
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Label,
		&i.Description,
		&i.StoreID,
	)
	return i, err
}

const createVariantOption = `-- name: CreateVariantOption :one
INSERT INTO variant_options (id, created_at, updated_at, variant_id, value, display_order, data, image_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING id, created_at, updated_at, variant_id, value, display_order, data, image_id
`

type CreateVariantOptionParams struct {
	ID           string           `json:"id"`
	CreatedAt    pgtype.Timestamp `json:"created_at"`
	UpdatedAt    pgtype.Timestamp `json:"updated_at"`
	VariantID    string           `json:"variant_id"`
	Value        string           `json:"value"`
	DisplayOrder int32            `json:"display_order"`
	Data         pgtype.Text      `json:"data"`
	ImageID      pgtype.Int4      `json:"image_id"`
}

func (q *Queries) CreateVariantOption(ctx context.Context, arg CreateVariantOptionParams) (VariantOption, error) {
	row := q.db.QueryRow(ctx, createVariantOption,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.VariantID,
		arg.Value,
		arg.DisplayOrder,
		arg.Data,
		arg.ImageID,
	)
	var i VariantOption
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.VariantID,
		&i.Value,
		&i.DisplayOrder,
		&i.Data,
		&i.ImageID,
	)
	return i, err
}

const deleteVariant = `-- name: DeleteVariant :exec
DELETE FROM variants
WHERE id = $1
`

func (q *Queries) DeleteVariant(ctx context.Context, id string) error {
	_, err := q.db.Exec(ctx, deleteVariant, id)
	return err
}

const deleteVariantOption = `-- name: DeleteVariantOption :exec
DELETE FROM variant_options
WHERE id = $1
`

func (q *Queries) DeleteVariantOption(ctx context.Context, id string) error {
	_, err := q.db.Exec(ctx, deleteVariantOption, id)
	return err
}

const getVariant = `-- name: GetVariant :one
SELECT 
    v.id, v.created_at, v.updated_at, v.name, v.label, v.description, v.store_id,
    COALESCE(
        json_agg(
            json_build_object(
                'id', vo.id,
                'created_at', vo.created_at,
                'updated_at', vo.updated_at,
                'variant_id', vo.variant_id,
                'value', vo.value,
                'data', vo.data,
                'image_id', vo.image_id,
                'display_order', vo.display_order
            ) ORDER BY vo.display_order
        ) FILTER (WHERE vo.id IS NOT NULL),
        '[]'::json
    ) AS options
FROM variants v
LEFT JOIN variant_options vo ON v.id = vo.variant_id
WHERE v.id = $1 AND v.store_id = $2
GROUP BY v.id
LIMIT 1
`

type GetVariantParams struct {
	ID      string `json:"id"`
	StoreID string `json:"store_id"`
}

type GetVariantRow struct {
	ID          string           `json:"id"`
	CreatedAt   pgtype.Timestamp `json:"created_at"`
	UpdatedAt   pgtype.Timestamp `json:"updated_at"`
	Name        string           `json:"name"`
	Label       string           `json:"label"`
	Description pgtype.Text      `json:"description"`
	StoreID     string           `json:"store_id"`
	Options     interface{}      `json:"options"`
}

func (q *Queries) GetVariant(ctx context.Context, arg GetVariantParams) (GetVariantRow, error) {
	row := q.db.QueryRow(ctx, getVariant, arg.ID, arg.StoreID)
	var i GetVariantRow
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Label,
		&i.Description,
		&i.StoreID,
		&i.Options,
	)
	return i, err
}

const getVariantAndOptionsArrayForVariantIds = `-- name: GetVariantAndOptionsArrayForVariantIds :many
SELECT
    v.id AS variant_id,
    COALESCE(
        json_agg(vo.id) FILTER (WHERE vo.id IS NOT NULL),
        '[]'::json
    )::json AS variant_options
FROM
    variants v
LEFT JOIN
    variant_options vo ON v.id = vo.variant_id
WHERE
    v.id = ANY($1::char(11)[])
    AND v.store_id = $2
GROUP BY
    v.id
ORDER BY
    v.created_at
`

type GetVariantAndOptionsArrayForVariantIdsParams struct {
	Column1 []string `json:"column_1"`
	StoreID string   `json:"store_id"`
}

type GetVariantAndOptionsArrayForVariantIdsRow struct {
	VariantID      string `json:"variant_id"`
	VariantOptions []byte `json:"variant_options"`
}

func (q *Queries) GetVariantAndOptionsArrayForVariantIds(ctx context.Context, arg GetVariantAndOptionsArrayForVariantIdsParams) ([]GetVariantAndOptionsArrayForVariantIdsRow, error) {
	rows, err := q.db.Query(ctx, getVariantAndOptionsArrayForVariantIds, arg.Column1, arg.StoreID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetVariantAndOptionsArrayForVariantIdsRow
	for rows.Next() {
		var i GetVariantAndOptionsArrayForVariantIdsRow
		if err := rows.Scan(&i.VariantID, &i.VariantOptions); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getVariantOption = `-- name: GetVariantOption :one
SELECT id, created_at, updated_at, variant_id, value, display_order, data, image_id FROM variant_options
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetVariantOption(ctx context.Context, id string) (VariantOption, error) {
	row := q.db.QueryRow(ctx, getVariantOption, id)
	var i VariantOption
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.VariantID,
		&i.Value,
		&i.DisplayOrder,
		&i.Data,
		&i.ImageID,
	)
	return i, err
}

const getVariantsWithOptionIds = `-- name: GetVariantsWithOptionIds :many
SELECT
    v.id AS variant_id,
    COALESCE(
        json_agg(
            json_build_object(
                'id', vo.id,
                'value', vo.value,
                'data', vo.data,
                'image_id', vo.image_id,
                'display_order', vo.display_order
            )
            ORDER BY vo.display_order
        ) FILTER (WHERE vo.id IS NOT NULL),
        '[]'::json
    )::json AS option_ids
FROM
    variants v
LEFT JOIN
    variant_options vo ON v.id = vo.variant_id
WHERE
    v.store_id = $1
GROUP BY
    v.id
ORDER BY
    v.created_at
`

type GetVariantsWithOptionIdsRow struct {
	VariantID string `json:"variant_id"`
	OptionIds []byte `json:"option_ids"`
}

func (q *Queries) GetVariantsWithOptionIds(ctx context.Context, storeID string) ([]GetVariantsWithOptionIdsRow, error) {
	rows, err := q.db.Query(ctx, getVariantsWithOptionIds, storeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetVariantsWithOptionIdsRow
	for rows.Next() {
		var i GetVariantsWithOptionIdsRow
		if err := rows.Scan(&i.VariantID, &i.OptionIds); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listVariantOptions = `-- name: ListVariantOptions :many
SELECT id, created_at, updated_at, variant_id, value, display_order, data, image_id FROM variant_options
WHERE variant_id = $1
ORDER BY display_order
LIMIT $2 OFFSET $3
`

type ListVariantOptionsParams struct {
	VariantID string `json:"variant_id"`
	Limit     int32  `json:"limit"`
	Offset    int32  `json:"offset"`
}

func (q *Queries) ListVariantOptions(ctx context.Context, arg ListVariantOptionsParams) ([]VariantOption, error) {
	rows, err := q.db.Query(ctx, listVariantOptions, arg.VariantID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []VariantOption
	for rows.Next() {
		var i VariantOption
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.VariantID,
			&i.Value,
			&i.DisplayOrder,
			&i.Data,
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

const listVariants = `-- name: ListVariants :many
SELECT 
    v.id, v.created_at, v.updated_at, v.name, v.label, v.description, v.store_id,
    COALESCE(
        json_agg(
            json_build_object(
                'id', vo.id,
                'created_at', vo.created_at,
                'updated_at', vo.updated_at,
                'variant_id', vo.variant_id,
                'value', vo.value,
                'data', vo.data,
                'image_id', vo.image_id,
                'display_order', vo.display_order
            ) ORDER BY vo.display_order
        ) FILTER (WHERE vo.id IS NOT NULL),
        '[]'::json
    ) AS options
FROM variants v
LEFT JOIN variant_options vo ON v.id = vo.variant_id
WHERE v.store_id = $1
GROUP BY v.id
ORDER BY v.created_at
LIMIT COALESCE($2, 10) OFFSET COALESCE($3, 0)
`

type ListVariantsParams struct {
	StoreID string      `json:"store_id"`
	Column2 interface{} `json:"column_2"`
	Column3 interface{} `json:"column_3"`
}

type ListVariantsRow struct {
	ID          string           `json:"id"`
	CreatedAt   pgtype.Timestamp `json:"created_at"`
	UpdatedAt   pgtype.Timestamp `json:"updated_at"`
	Name        string           `json:"name"`
	Label       string           `json:"label"`
	Description pgtype.Text      `json:"description"`
	StoreID     string           `json:"store_id"`
	Options     interface{}      `json:"options"`
}

func (q *Queries) ListVariants(ctx context.Context, arg ListVariantsParams) ([]ListVariantsRow, error) {
	rows, err := q.db.Query(ctx, listVariants, arg.StoreID, arg.Column2, arg.Column3)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListVariantsRow
	for rows.Next() {
		var i ListVariantsRow
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Name,
			&i.Label,
			&i.Description,
			&i.StoreID,
			&i.Options,
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

const listVariantsByIds = `-- name: ListVariantsByIds :many
SELECT 
    v.id, v.created_at, v.updated_at, v.name, v.label, v.description, v.store_id,
    COALESCE(
        json_agg(
            json_build_object(
                'id', vo.id,
                'created_at', vo.created_at,
                'updated_at', vo.updated_at,
                'variant_id', vo.variant_id,
                'value', vo.value,
                'data', vo.data,
                'image_id', vo.image_id,
                'display_order', vo.display_order
            ) ORDER BY vo.display_order
        ) FILTER (WHERE vo.id IS NOT NULL),
        '[]'::json
    )::json AS options
FROM variants v
LEFT JOIN variant_options vo ON v.id = vo.variant_id
WHERE v.id = ANY($1::char(11)[]) and v.store_id = $2
GROUP BY v.id
ORDER BY v.created_at
`

type ListVariantsByIdsParams struct {
	Column1 []string `json:"column_1"`
	StoreID string   `json:"store_id"`
}

type ListVariantsByIdsRow struct {
	ID          string           `json:"id"`
	CreatedAt   pgtype.Timestamp `json:"created_at"`
	UpdatedAt   pgtype.Timestamp `json:"updated_at"`
	Name        string           `json:"name"`
	Label       string           `json:"label"`
	Description pgtype.Text      `json:"description"`
	StoreID     string           `json:"store_id"`
	Options     []byte           `json:"options"`
}

func (q *Queries) ListVariantsByIds(ctx context.Context, arg ListVariantsByIdsParams) ([]ListVariantsByIdsRow, error) {
	rows, err := q.db.Query(ctx, listVariantsByIds, arg.Column1, arg.StoreID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListVariantsByIdsRow
	for rows.Next() {
		var i ListVariantsByIdsRow
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Name,
			&i.Label,
			&i.Description,
			&i.StoreID,
			&i.Options,
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

const updateVariant = `-- name: UpdateVariant :one
UPDATE variants
SET name = $2, description = $3, updated_at = $4, label = $5
WHERE id = $1
RETURNING id, created_at, updated_at, name, label, description, store_id
`

type UpdateVariantParams struct {
	ID          string           `json:"id"`
	Name        string           `json:"name"`
	Description pgtype.Text      `json:"description"`
	UpdatedAt   pgtype.Timestamp `json:"updated_at"`
	Label       string           `json:"label"`
}

func (q *Queries) UpdateVariant(ctx context.Context, arg UpdateVariantParams) (Variant, error) {
	row := q.db.QueryRow(ctx, updateVariant,
		arg.ID,
		arg.Name,
		arg.Description,
		arg.UpdatedAt,
		arg.Label,
	)
	var i Variant
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Label,
		&i.Description,
		&i.StoreID,
	)
	return i, err
}

const updateVariantOption = `-- name: UpdateVariantOption :one
UPDATE variant_options
SET value = $2, display_order = $3, updated_at = $4, image_id = $5, data = $6
WHERE id = $1
RETURNING id, created_at, updated_at, variant_id, value, display_order, data, image_id
`

type UpdateVariantOptionParams struct {
	ID           string           `json:"id"`
	Value        string           `json:"value"`
	DisplayOrder int32            `json:"display_order"`
	UpdatedAt    pgtype.Timestamp `json:"updated_at"`
	ImageID      pgtype.Int4      `json:"image_id"`
	Data         pgtype.Text      `json:"data"`
}

func (q *Queries) UpdateVariantOption(ctx context.Context, arg UpdateVariantOptionParams) (VariantOption, error) {
	row := q.db.QueryRow(ctx, updateVariantOption,
		arg.ID,
		arg.Value,
		arg.DisplayOrder,
		arg.UpdatedAt,
		arg.ImageID,
		arg.Data,
	)
	var i VariantOption
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.VariantID,
		&i.Value,
		&i.DisplayOrder,
		&i.Data,
		&i.ImageID,
	)
	return i, err
}
