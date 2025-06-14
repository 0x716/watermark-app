package service

import (
	"context"
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/0x716/watermark-app/internal/config"
	"github.com/0x716/watermark-app/internal/model"
	"github.com/0x716/watermark-app/internal/repository"
	"github.com/0x716/watermark-app/internal/utils"
	"github.com/disintegration/imaging"
)

type WatermarkService interface {
	SaveWatermarks(files []*multipart.FileHeader, outputPath string, perm os.FileMode) ([]string, error)
	UpdateWatermark(id int64, name string, width int64, height int64, opacity float32) error
	DeleteWatermark(id int64) error
	GetWatermarkById(id int64) (*model.Watermark, error)
	GetAllWatermark() ([]*model.Watermark, error)
}

type watermarkServiceImpl struct {
	repository repository.WatermarkRepository
}

func NewWatermarkService() WatermarkService {
	return &watermarkServiceImpl{
		repository: repository.NewWatermarkRepository(),
	}
}

func (wsi *watermarkServiceImpl) SaveWatermarks(files []*multipart.FileHeader, outputPath string, perm os.FileMode) ([]string, error) {
	// 判斷Watermark文件夾是否存在
	if err := os.Mkdir(outputPath, perm); os.IsNotExist(err) {
		return nil, fmt.Errorf("save watermark failed: the folder hasn't create")
	}

	// 返回文件名列表
	filenames := make([]string, 0, config.GlobalConfig.Constants.StringLength)

	// 創建Repository，用作操作Watermark表
	watermarkRepository := repository.NewWatermarkRepository()

	for _, file := range files {
		fh, err := file.Open()
		if err != nil {
			return nil, err
		}
		defer fh.Close()

		filePath, err := utils.SaveFile(fh, file.Filename, outputPath)
		if err != nil {
			return nil, err
		}

		// 存儲成功後，把圖片加載成img，然後獲取相應的Info
		img, err := imaging.Open(filePath)
		if err != nil {
			return nil, err
		}

		// 把Watermark的資料加載進數據庫中
		_, err = watermarkRepository.Create(context.Background(), file.Filename, int64(img.Bounds().Dx()), int64(img.Bounds().Dy()), 1.0)

		if err != nil {
			return nil, err
		}

		// 把filename加入filenames用作返回數據
		filenames = append(filenames, file.Filename)
	}

	return filenames, nil
}

func (wsi *watermarkServiceImpl) UpdateWatermark(id int64, name string, width int64, height int64, opacity float32) error {
	// 判斷傳入的值是否合理
	if id <= 0 || name == "" || width == 0 || height == 0 || opacity == 0.0 {
		return fmt.Errorf("invalid input")
	}

	return wsi.repository.Update(context.Background(), id, name, width, height, opacity)
}

func (wsi *watermarkServiceImpl) DeleteWatermark(id int64) error {
	// 判斷傳入的值是否合理
	if id <= 0 {
		return fmt.Errorf("invalid input")
	}

	watermark, err := wsi.repository.Get(context.Background(), id)
	if err != nil {
		return err
	}

	filePath := filepath.Join(config.GlobalConfig.Watermark.WatermarkDir, watermark.Name)

	err = os.Remove(filePath)
	if err != nil {
		return err
	}

	return wsi.repository.Delete(context.Background(), id)
}

func (wsi *watermarkServiceImpl) GetWatermarkById(id int64) (*model.Watermark, error) {
	// 判斷傳入的值是否合理
	if id <= 0 {
		return nil, fmt.Errorf("invalid input")
	}

	return wsi.repository.Get(context.Background(), id)
}

func (wsi *watermarkServiceImpl) GetAllWatermark() ([]*model.Watermark, error) {
	return wsi.repository.GetAll(context.Background())
}
