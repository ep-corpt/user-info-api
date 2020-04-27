package common

import (
	"encoding/json"
	"log"
	"net/http"
	"user-info-api/model"
)

const(
	statusSuccess = "00000"
	statusFailed = "00001"
)

func RespErr(w http.ResponseWriter, m string){
	w.Header().Set("Content-type", "application/json")
	if err:=json.NewEncoder(w).Encode(model.Response{Status: statusFailed, Desc: m}); err!=nil{
		log.Panicln(err)
	}
}

func RespSuccess(w http.ResponseWriter, i interface{}){
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(i)
}