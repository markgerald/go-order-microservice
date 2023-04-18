package message

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
	"os"
)

func Producer(orderId string) {
	w := &kafka.Writer{
		Addr:     kafka.TCP(os.Getenv("KAFKAHOST") + ":" + os.Getenv("KAFKAPORT")),
		Topic:    "go-gin-order",
		Balancer: &kafka.LeastBytes{},
	}

	err := w.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte("order_id"),
			Value: []byte(orderId),
		},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := w.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}
