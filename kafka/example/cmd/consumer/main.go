package main

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

func main() {
	// Writer برای ارسال پیام به Kafka
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:9092"}, // آدرس Kafka Broker که با compose اجرا کردیم
		Topic:   "test-topic",               // نام تاپیک
	})

	// ۱۰ تا پیام به تاپیک ارسال می کنیم
	for i := 1; i <= 10; i++ {
		msg := kafka.Message{
			Key:   []byte(fmt.Sprintf("Key-%d", i)),
			Value: []byte(fmt.Sprintf("Hello Kafka %d", i)),
		}
		if err := w.WriteMessages(context.Background(), msg); err != nil {
			log.Fatal("failed to write messages:", err)
		}
		log.Printf("Message %d sent!", i)
	}

	w.Close()
}
