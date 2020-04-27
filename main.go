package main

import (
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"net/http"
	"user-info-api/config"
	"user-info-api/service"
)

const (
	pathInquiry = "/inquiry"
)

func main() {
	config.InitConfig()

	db := config.InitDB()
	defer db.Close()

	c := config.InitConsumer()

	h := service.NewHandler(db, c)
	go h.Consumer()

	r := initRouter(h)
	panic(http.ListenAndServe(viper.GetString("server.port"), r))
}

func initRouter(h *service.Handler) *mux.Router{
	r := mux.NewRouter()
	r.HandleFunc(pathInquiry, h.Inquiry).Methods(http.MethodPost)
	return r
}