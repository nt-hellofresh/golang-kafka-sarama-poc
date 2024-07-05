package main

import (
	"net/http"
)

type Server struct {
	r    *http.ServeMux
	port string
}

func NewServer(port, topic string, brokerURLs ...string) *Server {
	return &Server{
		r:    configureRoutes(topic, brokerURLs...),
		port: port,
	}
}

func (s *Server) Start() error {
	return http.ListenAndServe(s.port, s.r)
}

func configureRoutes(topic string, brokerURLs ...string) *http.ServeMux {
	publisher := NewKafkaPublisher(topic, brokerURLs...)
	handler := NewRouteHandler(publisher)

	r := http.NewServeMux()
	r.HandleFunc("GET /", handler.index)
	r.HandleFunc("POST /foo", handler.foo)
	r.HandleFunc("POST /bar", handler.bar)
	return r
}
