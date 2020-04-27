package service

import (
	"encoding/json"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/jinzhu/gorm"
	"log"
	"user-info-api/entity"
	"user-info-api/model"
)

const (
	invalidRequest = "Invalid Request from kafka"
)

type consumerTask struct{
	db *gorm.DB
	c *kafka.Consumer
	rq *model.UserDetailWrapper
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

func (t *consumerTask)execute(m []byte) {
	t.initReq(m)
	t.save()
}

func (t *consumerTask)initTask(h *Handler){
	t.db = h.db
	t.c = h.c
}

func (t *consumerTask)initReq(m []byte){
	var w model.UserDetailWrapper
	if err:=json.Unmarshal(m, &w); err!=nil{
		log.Printf(invalidRequest)
		return
	}
	t.rq = &w
	log.Printf("Value is %v\n", t.rq.UserDetail)
}

func (t *consumerTask)save(){
	e := t.initEntity()
	rs := t.db.Create(&e)
	if rs.Error != nil {
		log.Printf("Exception occurred while save data %v", rs.Error)
	}
}

func(t *consumerTask)initEntity() *entity.UserInfo{
	var u entity.UserInfo
	u.Username = t.rq.CredentialDetail.Username
	u.FirstName = t.rq.UserDetail.FirstName
	u.LastName = t.rq.UserDetail.LastName
	return &u
}