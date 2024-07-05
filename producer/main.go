package main

import (
	"log/slog"
	"os"
)

func main() {
	s := NewServer(":2384", "foo-bar-topic", "localhost:29092")

	slog.Info("starting server",
		"URL", "http://localhost:2384",
		"kafkaTopic", "foo-bar-topic",
		"brokerURL", "localhost:29092",
	)

	if err := s.Start(); err != nil {
		slog.Error("could not start server: %v", err)
		os.Exit(1)
	}
}
