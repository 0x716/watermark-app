package utils

import (
	"fmt"
	"math/rand"
	"path/filepath"
	"time"

	"github.com/0x716/watermark-app/internal/config"
)

func GenerateFilename(name string) string {
	timestramp := time.Now().Format("20060102_150405")
	rng := rand.New(rand.NewSource(int64(config.GlobalConfig.Image.RandomSeed) + time.Now().UnixNano()))
	randomNum := rng.Intn(int(config.GlobalConfig.Constants.RandomLimit))

	ext := filepath.Ext(name)
	filename := fmt.Sprintf("%s_%06d%s", timestramp, randomNum, ext)

	return filename
}
