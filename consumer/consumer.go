package main

import (
	"bytes"
	"github.com/IBM/sarama"
	"kafka_sarama/pkg"
	"log"
)

// Consumer represents a Sarama consumer group consumer
type Consumer struct {
	ready   chan bool
	process MessageHandler
}

func NewConsumer(handler MessageHandler) *Consumer {
	return &Consumer{
		ready:   make(chan bool),
		process: handler,
	}
}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (consumer *Consumer) Setup(sarama.ConsumerGroupSession) error {
	// Mark the consumer as ready
	close(consumer.ready)
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (consumer *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
// Once the Messages() channel is closed, the Handler must finish its processing
// loop and exit.
func (consumer *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	// NOTE:
	// Do not move the code below to a goroutine.
	// The `ConsumeClaim` itself is called within a goroutine, see:
	// https://github.com/IBM/sarama/blob/main/consumer_group.go#L27-L29
	for {
		select {
		case message, ok := <-claim.Messages():
			if !ok {
				log.Printf("message channel was closed")
				return nil
			}

			msg := new(pkg.Message)
			buf := bytes.NewBuffer(message.Value)
			if err := pkg.Decode(buf, msg); err != nil {
				log.Printf("could not decode message: %v", err)
				continue
			}

			if err := consumer.process(msg); err != nil {
				log.Printf("could not process message: %v", err)
				continue
			}

			session.MarkMessage(message, "")
		// Should return when `session.Context()` is done.
		// If not, will raise `ErrRebalanceInProgress` or `read tcp <ip>:<port>: i/o timeout` when kafka rebalance. see:
		// https://github.com/IBM/sarama/issues/1192
		case <-session.Context().Done():
			return nil
		}
	}
}
