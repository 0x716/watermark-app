-- name: CreateWatermark :one
INSERT INTO watermark (name, width, height, opacity) VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetWatermark :one
SELECT * FROM watermark WHERE id = $1;

-- name: ListWatermark :many
SELECT * FROM watermark;

-- name: UpdateWatermark :exec
UPDATE watermark SET
name = $2,
width = $3,
height = $4,
opacity = $5,
update_at = NOW()
WHERE id = $1;

-- name: DeleteWatermark :exec
DELETE FROM watermark WHERE id = $1;
