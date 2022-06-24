package controllers

import (
	"context"
	"fish_go_api/config"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"cloud.google.com/go/storage"
	"github.com/gofiber/fiber/v2"
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

	image, err := file.Open()
	if err != nil {
		log.Printf("failed to open image: %s", err)
		return c.JSON(fiber.Map{
			"message": "failed to open image",
		})
	}

	log.Println("start upload to GCS")

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
	if _, err := io.Copy(wc, image); err != nil {
		return c.JSON(fiber.Map{
			"message": fmt.Sprintf("io.Copy: %v", err),
		})
	}
	if err := wc.Close(); err != nil {
		return c.JSON(fiber.Map{
			"message": fmt.Sprintf("Writer.Close: %v", err),
		})
	}

	log.Println("Success to upload image")
	return c.JSON(fiber.Map{
		"url": objectPath + filename,
	})
}
