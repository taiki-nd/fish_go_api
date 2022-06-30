package controllerlogics

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"mime/multipart"
	"os"

	"github.com/nfnt/resize"
)

/*
	resizeImg
	resize image size
	@params img image.Image
	@params fileData multipart.File
	@params filename string
	@return string
	@return error
*/
func ResizeImg(img image.Image, fileData multipart.File, data, filename string) (string, error) {
	log.Println("start to resize image")
	const width = 800
	const height = 0
	resizedImg := resize.Resize(width, height, img, resize.NearestNeighbor)

	osPath := "uploads/" + filename
	output, err := os.Create(osPath)
	if err != nil {
		log.Printf("failed to create %v: %v", osPath, err)
		return "", fmt.Errorf("failed to create %v: %v", osPath, err)
	}

	switch data {
	case "png":
		err := png.Encode(output, resizedImg)
		if err != nil {
			log.Printf("failed to encode image: %v", err)
			return "", fmt.Errorf("failed to encode image: %v", err)
		}
	case "jpeg", "jpg":
		opts := &jpeg.Options{Quality: 100}
		err := jpeg.Encode(output, resizedImg, opts)
		if err != nil {
			log.Printf("failed to encode image: %v", err)
			return "", fmt.Errorf("failed to encode image: %v", err)
		}
	default:
		err := png.Encode(output, resizedImg)
		if err != nil {
			log.Printf("failed to encode image: %v", err)
			return "", fmt.Errorf("failed to encode image: %v", err)
		}
	}

	return osPath, nil
}
