package service

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/0x716/watermark-app/internal/service"
	"github.com/stretchr/testify/assert"
)

func TestAddwatermark(t *testing.T) {
	curDir, _ := os.Getwd()
	baseImagePath := filepath.Join(filepath.Join(curDir, "./inputs/image.png"))
	watermarkPath := filepath.Join("./logo/", "logo.png")
	ouputPath := filepath.Join("./outputs", "output.png")

	err := service.AddWatermark(baseImagePath, watermarkPath, ouputPath, 0.3, 0.5)
	assert.NoError(t, err)
}
