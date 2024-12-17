package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

const (
	imageDir    = "./images"
	videoDir    = "./videos"
	documentDir = "./documents"
	audioDir    = "./audios"
)

var directoriesByFile = map[string]string{
	"image":    imageDir,
	"video":    videoDir,
	"document": documentDir,
	"audio":    audioDir,
}

func HandleFileCreated(filePath string) error {
	fileType := getFileType(filePath)
	destPath, ok := directoriesByFile[fileType]
	if !ok {
		fmt.Printf("Warning: Unknown file type for %s\n", filePath)
		return nil
	}

	err := MoveFile(filePath, destPath)
	if err != nil {
		return fmt.Errorf("error moving file: %v", err)
	}
	return nil
}

func MoveFile(sourcePath, destDir string) error {
	// Ensure the destination directory exists
	err := os.MkdirAll(destDir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("couldn't create destination directory: %v", err)
	}

	filename := filepath.Base(sourcePath)
	destPath := filepath.Join(destDir, filename)

	inputFile, err := os.Open(sourcePath)
	if err != nil {
		return fmt.Errorf("couldn't open source file: %v", err)
	}
	defer inputFile.Close()

	outputFile, err := os.Create(destPath)
	if err != nil {
		return fmt.Errorf("couldn't create destination file: %v", err)
	}
	defer outputFile.Close()

	_, err = io.Copy(outputFile, inputFile)
	if err != nil {
		return fmt.Errorf("couldn't copy to destination file: %v", err)
	}

	if err = os.Remove(sourcePath); err != nil {
		return fmt.Errorf("couldn't remove source file: %v", err)
	}
	return nil
}

func getFileType(filePath string) string {
	imageExtensions := map[string]bool{".jpg": true, ".jpeg": true, ".png": true, ".gif": true, ".bmp": true, ".svg": true, ".webp": true}
	videoExtensions := map[string]bool{".mp4": true, ".mov": true, ".avi": true, ".mkv": true, ".flv": true, ".wmv": true}
	documentExtensions := map[string]bool{".pdf": true, ".doc": true, ".docx": true, ".xls": true, ".xlsx": true, ".txt": true, ".ppt": true, ".pptx": true}
	audioExtensions := map[string]bool{".mp3": true, ".wav": true, ".aac": true, ".flac": true, ".ogg": true}

	ext := strings.ToLower(filepath.Ext(filePath))

	switch {
	case imageExtensions[ext]:
		return "image"
	case videoExtensions[ext]:
		return "video"
	case documentExtensions[ext]:
		return "document"
	case audioExtensions[ext]:
		return "audio"
	default:
		return "unknown"
	}
}
