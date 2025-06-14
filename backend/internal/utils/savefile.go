package utils

import (
	"io"
	"os"
	"path/filepath"
)

func SaveFile(srcFile io.Reader, filename string, outputDir string) (string, error) {
	outputPath := filepath.Join(outputDir, filename)

	dstFile, err := os.Create(outputPath)
	if err != nil {
		return "", err
	}

	_, err = io.Copy(dstFile, srcFile)

	return outputPath, err
}
