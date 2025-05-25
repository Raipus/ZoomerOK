package config

import (
	"fmt"
	"log"
	"os"
)

func readImageToBytes(imagePath string) ([]byte, error) {
	file, err := os.Open(imagePath)
	if err != nil {
		return nil, fmt.Errorf("ошибка при открытии файла: %w", err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Printf("Ошибка при закрытии файла: %v", err)
		}
	}()

	imageData, err := os.ReadFile(imagePath)
	if err != nil {
		return nil, fmt.Errorf("ошибка при чтении файла: %w", err)
	}

	return imageData, nil
}

func getByteImage(c *ConfigType) []byte {
	imageData, err := readImageToBytes(c.Photo.Image)
	if err != nil {
		log.Fatalf("Ошибка: %s", err)
	}

	return imageData
}
