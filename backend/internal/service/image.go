package service

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/0x716/watermark-app/internal/config"
	"github.com/0x716/watermark-app/internal/utils"
	"github.com/disintegration/imaging"
)

type UploadResult struct {
	Name string `json:"name"`
	Size int64  `json:"size"`
	URL  string `json:"url"`
}

type Image interface {
	SaveUploadImage(file *multipart.FileHeader, outputDir string, perm os.FileMode) (string, error)
}

func SaveUploadImages(files []*multipart.FileHeader, outputDir string, perm os.FileMode) ([]string, error) {
	filepaths := make([]string, 0, config.GlobalConfig.Constants.StringLength)
	for _, file := range files {
		fd, err := file.Open()
		if err != nil {
			return nil, err
		}
		defer fd.Close()

		filename, err := utils.SaveFile(fd, file.Filename, config.GlobalConfig.Image.ImageDir)
		if err != nil {
			return nil, err
		}

		filePath := filepath.Join(config.GlobalConfig.Image.ImageDir, filename)

		filepaths = append(filepaths, filePath)
	}

	return filepaths, nil
}

func AddWatermark(imagePath, watermarkPath, outputPath string, opacity float32, scale float32) error {
	// Load upload image
	img, err := imaging.Open(imagePath)
	if err != nil {
		return fmt.Errorf("failed to load watermark: %w", err)
	}

	// Load watermark
	watermark, err := imaging.Open(watermarkPath)
	if err != nil {
		return fmt.Errorf("failed to load watermark: %w", err)
	}

	// ✅ Resize watermark before applying opacity
	newWidth := int(float32(watermark.Bounds().Dx()) * scale)
	newHeight := int(float32(watermark.Bounds().Dy()) * scale)
	watermark = imaging.Resize(watermark, newWidth, newHeight, imaging.Lanczos)

	// ✅ Then apply opacity
	watermark = adjustOpacity(watermark, opacity)

	// Watermark Position
	offsetX := img.Bounds().Dx() - watermark.Bounds().Dx() - 10
	offsetY := img.Bounds().Dy() - watermark.Bounds().Dy() - 10
	offset := image.Pt(offsetX, offsetY)

	// New Canvas
	output := image.NewNRGBA(img.Bounds())
	draw.Draw(output, img.Bounds(), img, image.Point{}, draw.Over)
	draw.Draw(output, watermark.Bounds().Add(offset), watermark, image.Point{}, draw.Over)

	// Save new image
	err = imaging.Save(output, outputPath)
	if err != nil {
		return fmt.Errorf("failed to save watermarked image: %w", err)
	}

	return nil
}

func adjustOpacity(img image.Image, alpha float32) *image.NRGBA {
	bounds := img.Bounds()
	result := image.NewNRGBA(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			original := color.NRGBAModel.Convert(img.At(x, y)).(color.NRGBA)
			original.A = uint8(float32(original.A) * alpha)
			result.SetNRGBA(x, y, original)
		}
	}

	return result
}
