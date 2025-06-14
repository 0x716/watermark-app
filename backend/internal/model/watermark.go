package model

import (
	"path/filepath"
	"time"

	"github.com/0x716/watermark-app/internal/config"
	"github.com/0x716/watermark-app/internal/db"
)

type Watermark struct {
	Id       int64     `json:"id"`
	Name     string    `json:"name"`
	Width    int64     `json:"width"`
	Height   int64     `json:"height"`
	Opacity  float32   `json:"opacity"`
	CreateAt time.Time `json:"create_at"`
	Path     string    `json:"path"`
}

func WatermarkFrom(watermark db.Watermark) *Watermark {
	return &Watermark{
		Id:       watermark.ID,
		Name:     watermark.Name,
		Width:    watermark.Width,
		Height:   watermark.Height,
		Opacity:  watermark.Opacity,
		CreateAt: watermark.CreateAt,
		Path:     filepath.Join(config.GlobalConfig.Watermark.WatermarkDir, watermark.Name),
	}
}
