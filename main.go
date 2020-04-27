package main

import (
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"net/http"
	"user-info-api/config"
	"user-info-api/service"
)

func main() {
	config.InitConfig()

	db := config.InitDB()
	defer db.Close()

	c := config.InitConsumer()

	h := service.NewHandler(db, c)
	go h.Consumer()
	panic(http.ListenAndServe(viper.GetString("server.port"), mux.NewRouter()))

}