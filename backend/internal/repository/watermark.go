package repository

import (
	"context"

	"github.com/0x716/watermark-app/internal/db"
	"github.com/0x716/watermark-app/internal/infra"
	"github.com/0x716/watermark-app/internal/model"
)

type WatermarkRepository interface {
	Create(ctx context.Context, name string, width int64, height int64, opacity float32) (*model.Watermark, error)
	Get(ctx context.Context, id int64) (*model.Watermark, error)
	GetAll(ctx context.Context) ([]*model.Watermark, error)
	Update(ctx context.Context, id int64, name string, width int64, height int64, opacity float32) error
	Delete(ctx context.Context, id int64) error
}

type watermarkRepositoryImpl struct {
	Queries *db.Queries
}

func NewWatermarkRepository() WatermarkRepository {
	return &watermarkRepositoryImpl{
		Queries: db.New(infra.DB),
	}
}

func (wr *watermarkRepositoryImpl) Create(ctx context.Context, name string, width int64, height int64, opacity float32) (*model.Watermark, error) {
	dbWatermark, err := wr.Queries.CreateWatermark(ctx, db.CreateWatermarkParams{
		Name:    name,
		Width:   width,
		Height:  height,
		Opacity: opacity,
	})

	if err != nil {
		return nil, err
	}

	return model.WatermarkFrom(dbWatermark), nil
}

func (wr *watermarkRepositoryImpl) Get(ctx context.Context, id int64) (*model.Watermark, error) {
	dbWatermark, err := wr.Queries.GetWatermark(ctx, id)
	if err != nil {
		return nil, err
	}

	return model.WatermarkFrom(dbWatermark), nil
}

func (wr *watermarkRepositoryImpl) GetAll(ctx context.Context) ([]*model.Watermark, error) {
	watermarks, err := wr.Queries.ListWatermark(ctx)
	if err != nil {
		return nil, err
	}

	watermarkModels := make([]*model.Watermark, 0, 1000)

	for _, watermark := range watermarks {
		watermarkModels = append(watermarkModels, model.WatermarkFrom(watermark))
	}

	return watermarkModels, nil
}

func (wr *watermarkRepositoryImpl) Update(ctx context.Context, id int64, name string, width int64, height int64, opacity float32) error {
	err := wr.Queries.UpdateWatermark(ctx, db.UpdateWatermarkParams{
		ID:      id,
		Name:    name,
		Width:   width,
		Height:  height,
		Opacity: opacity,
	})

	return err
}

func (wr *watermarkRepositoryImpl) Delete(ctx context.Context, id int64) error {
	err := wr.Queries.DeleteWatermark(ctx, id)
	return err
}
