package main

import (
	"user-info-api/config"
	"user-info-api/service"
)

func main() {
	config.InitConfig()

	db := config.InitDB()
	defer db.Close()

	c := config.InitConsumer()

	h := service.NewHandler(db, c)
	h.Consumer()
}