package pkg

import "time"

const (
	FooEventType = "foo"
	BarEventType = "bar"
)

type Message struct {
	ID          string
	MessageType string
	At          time.Time
	Body        []byte
}

type FooContent struct {
	Title     string
	FirstName string
	LastName  string
}

type BarContent struct {
	Address  string
	PostCode string
	Region   string
}
