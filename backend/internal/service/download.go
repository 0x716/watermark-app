package service

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func CreateZipFromFiles(files []string, outputDirPath string, w io.Writer) error {
	buf := new(bytes.Buffer)
	zipWriter := zip.NewWriter(buf)

	for _, file := range files {
		filePath := filepath.Join(outputDirPath, file)
		err := addFileToZip(zipWriter, filePath)
		if err != nil {
			return fmt.Errorf("add file %s to zip failed: %w", file, err)
		}
	}

	if err := zipWriter.Close(); err != nil {
		return fmt.Errorf("zipWriter close failed: %w", err)
	}

	if _, err := w.Write(buf.Bytes()); err != nil {
		return fmt.Errorf("zip close failed: %w", err)
	}

	return nil
}

func addFileToZip(zipWriter *zip.Writer, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return err
	}

	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}

	header.Name = filepath.Base(filePath)
	header.Method = zip.Deflate // 指定壓縮方法
	header.SetMode(info.Mode()) // 設定檔案權限
	header.Modified = info.ModTime()

	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}

	_, err = io.Copy(writer, file)

	return err
}
