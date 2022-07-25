package adaptor

import (
	"fmt"
	"log"
	"photorage-community/config"
)

// TODO : Kafka에 해당 Topic 생성 필요
const COMMENT_DELETE_MESSAGE_TOPIC string = "comment-delete-event"

func CommentDeleteConsumer() {
	kafkaConnection := config.Kafka()

	kafkaConnection.SubscribeTopics([]string{COMMENT_DELETE_MESSAGE_TOPIC}, nil)
	defer func() {
		log.Fatal("Kafka Subscribe Topics Failed.")
		kafkaConnection.Close()
	}()

	for {
		msg, err := kafkaConnection.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else {
			// The client will automatically try to recover from all errors.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}
