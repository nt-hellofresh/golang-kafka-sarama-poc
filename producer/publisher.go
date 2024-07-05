package main

import (
	"bytes"
	"github.com/IBM/sarama"
	"kafka_sarama/pkg"
	"log"
)

type Publisher interface {
	Publish(msg pkg.Message) error
}

type KafkaPublisher struct {
	brokerURLs []string
	topic      string
	conn       sarama.SyncProducer
}

func NewKafkaPublisher(topic string, brokerURLs ...string) *KafkaPublisher {
	cfg := sarama.NewConfig()
	cfg.Producer.RequiredAcks = sarama.WaitForAll
	cfg.Producer.Retry.Max = 5
	cfg.Producer.Return.Successes = true

	conn, err := sarama.NewSyncProducer(brokerURLs, cfg)

	if err != nil {
		log.Fatal(err)
	}

	return &KafkaPublisher{
		topic:      topic,
		brokerURLs: brokerURLs,
		conn:       conn,
	}
}

func (p *KafkaPublisher) Publish(msg pkg.Message) error {
	buf := new(bytes.Buffer)
	if err := pkg.Encode(buf, &msg); err != nil {
		return err
	}

	_, _, err := p.conn.SendMessage(&sarama.ProducerMessage{
		Topic: p.topic,
		Key:   sarama.ByteEncoder(msg.ID),
		Value: sarama.StringEncoder(buf.String()),
	})
	return err
}
