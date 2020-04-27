package service

import (
	"encoding/json"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/jinzhu/gorm"
	"log"
	"user-info-api/model"
)

const (
	invalidRequest = "Invalid Request from kafka"
)

type consumerTask struct{
	db *gorm.DB
	c *kafka.Consumer
	rq *model.UserDetail
}

func (h *Handler)Consumer(){
	var t consumerTask
	t.initTask(h)
	for {
		m, err := t.c.ReadMessage(-1)
		if err == nil{
			log.Printf("Message on %s: %s\n", m.TopicPartition, m.Value)
			t.execute(m.Value)
		}else{
			log.Printf("Consumer err : %v msg: %v\n", err, m)
		}
	}

}

func (t *consumerTask)initTask(h *Handler){
	t.db = h.db
	t.c = h.c
}

func (t *consumerTask)execute(m []byte) {
	var w model.UserDetailWrapper
	if err:=json.Unmarshal(m, &w); err!=nil{
		log.Printf(invalidRequest)
		return
	}
	t.rq = &w.UserDetail
	log.Printf("Value is %v\n", t.rq)
}
