// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0

package db

import (
	"time"
)

type Watermark struct {
	ID       int64
	Name     string
	Width    int64
	Height   int64
	Opacity  float32
	CreateAt time.Time
	UpdateAt time.Time
}
