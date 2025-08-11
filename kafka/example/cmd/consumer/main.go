package main

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

func main() {
	// Reader برای دریافت پیام از Kafka
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"}, // آدرس بروکر
		Topic:   "test-topic",               // نام تاپیک
		GroupID: "my-group",                 // نام گروه کانسیومر
	})

	// حلقه برای دریافت پیام ها
	for {
		// خواندن یک پیام
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Fatal("failed to read message:", err)
		}
		log.Printf("Message received: key=%s value=%s\n", string(m.Key), string(m.Value))
	}
}
