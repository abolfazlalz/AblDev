# Kafka Example with Go

This project is a simple example of using **Apache Kafka** with **Go**.
It demonstrates how to implement a **Producer** to send messages and a **Consumer** to receive messages, while running Kafka locally using **Docker Compose**.

## 📂 Project Structure

```
kafka/example/
├── cmd/
│ ├── consumer/ # Kafka message consumer
│ └── producer/ # Kafka message producer
├── compose.yml # Docker Compose configuration for Kafka and Zookeeper
├── go.mod
└── go.sum
```

## 🔹 What This Example Does
1. Runs a Kafka broker and Zookeeper locally using Docker Compose.
2. **Producer** sends multiple sample messages to a specific Kafka topic.
3. **Consumer** listens to the same topic, receives the messages, and logs them to the console.
4. This demonstrates the complete flow of producing and consuming messages in Kafka.

## 🚀 How to Run

1. Start Kafka and Zookeeper:

```bash
docker compose -f compose.yml up -d
```

2. Run the producer:

```bash
go run cmd/producer/main.go
```

3. Run the consumer:

```bash
go run cmd/consumer/main.go
```

4. Verify the messages are received by the consumer.

5. Stop Kafka and Zookeeper:

```bash
docker compose -f compose.yml down
```

## 🛠 Requirements

- Go 1.20+
- Docker & Docker Compose
