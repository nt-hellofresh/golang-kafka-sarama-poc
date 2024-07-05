package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"kafka_sarama/pkg"
	"log"
)

type MessageHandler func(msg *pkg.Message) error

func HandleMessage(msg *pkg.Message) error {
	log.Printf("received message: %v\n", msg)
	switch msg.MessageType {
	case pkg.FooEventType:
		return handleFoo(msg)
	case pkg.BarEventType:
		return handleBar(msg)
	default:
		return handleDefault(msg)
	}
}

func handleFoo(msg *pkg.Message) error {
	buf := bytes.NewBuffer(msg.Body)
	fooContent := new(pkg.FooContent)

	if err := json.NewDecoder(buf).Decode(fooContent); err != nil {
		return err
	}

	log.Printf("foo content: %v\n", fooContent)
	return nil
}

func handleBar(msg *pkg.Message) error {
	buf := bytes.NewBuffer(msg.Body)
	barContent := new(pkg.BarContent)

	if err := json.NewDecoder(buf).Decode(barContent); err != nil {
		return err
	}

	log.Printf("bar content: %v\n", barContent)
	return nil
}

func handleDefault(_ *pkg.Message) error {
	return errors.New("unsupported message type")
}
