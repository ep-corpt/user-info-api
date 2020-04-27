package config

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
	"log"
)

const(
	configName = "config"
	dot = "."
)

func InitConfig(){
	viper.SetConfigName(configName)
	viper.AddConfigPath(dot)
	if err:=viper.ReadInConfig();err!= nil {
		log.Panicln(err)
	}
}

func InitDB() *gorm.DB{
	db, err := gorm.Open("postgres", viper.GetString("db"))
	if err!= nil {
		log.Panicln(err)
	}
	return db
}

func InitConsumer() *kafka.Consumer{
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": viper.GetString("kafka.host"),
		"group.id":         viper.GetString("kafka.group"),
		"auto.offset.reset": viper.GetString("kafka.auto-offset"),
	})

	if err!= nil {
		log.Panicln(err)
	}

	if err:=c.SubscribeTopics([]string{viper.GetString("kafka.topic"), "^aRegex.*[Tt]opic"}, nil);err!=nil{
		log.Panicln(err)
	}
	return c
}