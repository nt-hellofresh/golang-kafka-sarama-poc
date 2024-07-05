package pkg

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
	"io"
	"kafka_sarama/pkg/pb"
)

func Encode(w io.Writer, msg *Message) error {
	payload := &pb.Message{
		ID:          msg.ID,
		MessageType: msg.MessageType,
		At:          timestamppb.New(msg.At),
		Body:        msg.Body,
	}

	data, err := proto.Marshal(payload)
	if err != nil {
		return err
	}

	_, err = w.Write(data)
	return err
}

func Decode(r io.Reader, msg *Message) error {
	data, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	payload := new(pb.Message)
	if err := proto.Unmarshal(data, payload); err != nil {
		return err
	}

	msg.ID = payload.GetID()
	msg.MessageType = payload.GetMessageType()
	msg.At = payload.GetAt().AsTime()
	msg.Body = payload.GetBody()

	return nil
}
