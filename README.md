# Golang-Kafka-Sarama-POC

This is a simple example of how to use the Sarama Kafka client in Golang for both producing and consuming messages.

## Prerequisites
- [Golang](https://go.dev/)
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

## Setup Instructions
1. Clone the repository.
2. Run `make dev-up` to bring up the local Kafka cluster.
3. Run main.go in the producer folder.
4. Open the http://localhost:9021/ to see a web interface for producing Kafka messages.
5. Run main.go in the consumer folder to consume messages from the topic.
