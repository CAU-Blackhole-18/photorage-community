package config

import (
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

const KAFKA_SERVER_INFO string = "localhost:9092"
const KAFKA_GROUP_ID string = "community.group"
const KAFKA_AUTO_OFFSET_RESET_OPTION string = "earliest" // Default : Latest

func Kafka() *kafka.Consumer {

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": KAFKA_SERVER_INFO,
		"group.id":          KAFKA_GROUP_ID,
		"auto.offset.reset": KAFKA_AUTO_OFFSET_RESET_OPTION,
	})

	if err != nil {
		// error가 있으면 panic을 발생. defer로 예외처리 가능하긴 함
		panic(err)
	}

	return c
}
