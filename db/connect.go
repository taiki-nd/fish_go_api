package db

import (
	"fish_go_api/config"
	"fish_go_api/models"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True",
		config.Config.UserDevelop, config.Config.PasswordDevelop,
		config.Config.HostDevelop, config.Config.PortDevelop,
		config.Config.NameDevelop)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %s", err)
	}

	DB = db

	log.Printf("success db connection: %v", db)

	db.AutoMigrate(
		&models.User{},
		&models.CommentReply{},
		&models.GroundComment{},
		&models.Style{},
		&models.Howto{},
		&models.Fish{},
		&models.Ground{},
		&models.PostComment{},
		&models.Post{},
	)
}
