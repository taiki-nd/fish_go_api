package main

import (
	"fish_go_api/config"
	"fish_go_api/utils"

	"log"
)

func main() {
	utils.Logging(config.Config.Logfile)
	log.Println("hello world!")
}
