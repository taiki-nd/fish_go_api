package controllers

import (
	"context"
	"fish_go_api/config"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"os"
	"time"

	"cloud.google.com/go/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/nfnt/resize"
	"google.golang.org/api/option"
)

func ImageUpload(c *fiber.Ctx) error {
	log.Println("start to upload image")
	file, err := c.FormFile("image")
	if err != nil {
		log.Printf("uploads images error: %s", err)
		return c.JSON(fiber.Map{
			"message": "failed to upload image",
		})
	}

	filename := file.Filename

	fileData, err := file.Open()
	if err != nil {
		log.Printf("failed to open image: %s", err)
		return c.JSON(fiber.Map{
			"message": "failed to open image",
		})
	}

	// 画像をimage.Image型にdecode
	img, data, err := image.Decode(fileData)
	if err != nil {
		log.Printf("failed to decode image: %v", err)
		return c.JSON(fiber.Map{
			"message": fmt.Sprintf("failed to decode image: %v", err),
		})
	}
	fileData.Close()

	fmt.Println("width:", img.Bounds().Dx())
	fmt.Println("height:", img.Bounds().Dy())

	if img.Bounds().Dx() > 800 {
		log.Println("start to resize image")
		const width = 800
		const height = 0
		resizedImg := resize.Resize(width, height, img, resize.NearestNeighbor)

		osPath := "uploads/" + filename
		output, err := os.Create(osPath)
		if err != nil {
			log.Printf("failed to create %v: %v", osPath, err)
			return c.JSON(fiber.Map{
				"message": fmt.Sprintf("failed to create %v: %v", osPath, err),
			})
		}

		switch data {
		case "png":
			err := png.Encode(output, resizedImg)
			if err != nil {
				log.Printf("failed to encode image: %v", err)
				return c.JSON(fiber.Map{
					"message": fmt.Sprintf("failed to encode image: %v", err),
				})
			}
		case "jpeg", "jpg":
			opts := &jpeg.Options{Quality: 100}
			err := jpeg.Encode(output, resizedImg, opts)
			if err != nil {
				log.Printf("failed to encode image: %v", err)
				return c.JSON(fiber.Map{
					"message": fmt.Sprintf("failed to encode image: %v", err),
				})
			}
		default:
			err := png.Encode(output, resizedImg)
			if err != nil {
				log.Printf("failed to encode image: %v", err)
				return c.JSON(fiber.Map{
					"message": fmt.Sprintf("failed to encode image: %v", err),
				})
			}
		}
		fileData, err = os.Open(osPath)
		if err != nil {
			log.Printf("failed to open %v: %v", osPath, err)
			return c.JSON(fiber.Map{
				"message": fmt.Sprintf("failed to open %v: %v", osPath, err),
			})
		}
	}

	log.Println("start to upload image to GCS")

	jsonPath := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(jsonPath))
	if err != nil {
		log.Printf("failed to create client: %s", err)
		return c.JSON(fiber.Map{
			"message": fmt.Sprintf("failed to create client: %s", err),
		})
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	bucketName := config.Config.GcsBucketNameLocal
	objectPath := config.Config.GcsObjectPathLocal

	o := client.Bucket(bucketName).Object(filename)
	o = o.If(storage.Conditions{DoesNotExist: true})
	wc := o.NewWriter(ctx)
	if _, err := io.Copy(wc, fileData); err != nil {
		return c.JSON(fiber.Map{
			"message": fmt.Sprintf("io.Copy: %v", err),
		})
	}
	if err := wc.Close(); err != nil {
		return c.JSON(fiber.Map{
			"message": fmt.Sprintf("Writer.Close: %v", err),
		})
	}

	err = os.Remove("uploads/" + filename)
	if err != nil {
		log.Printf("failed to remove uploads/%v: %v", filename, err)
		return c.JSON(fiber.Map{
			"message": fmt.Sprintf("failed to remove uploads/%v: %v", filename, err),
		})
	}

	log.Println("Success to upload image")
	return c.JSON(fiber.Map{
		"url":      objectPath + filename,
		"filename": filename,
	})
}

func ImageDelete(filename string) string {
	log.Printf("start to delete image form GCS: %v", filename)

	jsonPath := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(jsonPath))
	if err != nil {
		log.Printf("failed to create client: %s", err)
		return fmt.Sprintf("failed to create client: %s", err)
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	bucketName := config.Config.GcsBucketNameLocal

	o := client.Bucket(bucketName).Object(filename)

	attrs, err := o.Attrs(ctx)
	if err != nil {
		log.Printf("object.Attrs: %v", err)
		return fmt.Sprintf("object.Attrs: %v", err)
	}
	o = o.If(storage.Conditions{GenerationMatch: attrs.Generation})

	err = o.Delete(ctx)
	if err != nil {
		log.Printf("Object.Delete: %v", err)
		return fmt.Sprintf("Object.Delete: %v", err)
	}

	return "success to delete image"
}
