package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"kafka_sarama/pkg"
	"log/slog"
	"net/http"
	"time"
)

type RouteHandler struct {
	publisher Publisher
	templates *template.Template
}

func NewRouteHandler(publisher Publisher) *RouteHandler {
	tmpl, err := template.New("").ParseGlob("producer/templates/*.html")

	if err != nil {
		panic(err)
	}

	return &RouteHandler{
		publisher: publisher,
		templates: tmpl,
	}
}

func (h *RouteHandler) index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	if err := h.templates.ExecuteTemplate(w, "index.html", nil); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "could not render home page: %v", err)
		return
	}
}

func (h *RouteHandler) foo(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	msg, err := createFooMessage()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "could not marshal foo content: %v", err)
		return
	}

	if err := h.publisher.Publish(*msg); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "could not publish message: %v", err)
	}

	logMsg := "successfully published the following message"
	slog.Info(logMsg, "topic", "foo-bar-topic", "message", msg)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"log":     logMsg,
		"payload": msg,
	})
}

func (h *RouteHandler) bar(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	msg, err := createBarMessage()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "could not marshal foo content: %v", err)
		return
	}
	if err := h.publisher.Publish(*msg); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "could not publish message: %v", err)
	}

	logMsg := "successfully published the following message"
	slog.Info(logMsg, "topic", "foo-bar-topic", "message", msg)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"log":     logMsg,
		"payload": msg,
	})
}

func createFooMessage() (*pkg.Message, error) {
	data, err := json.Marshal(pkg.FooContent{
		Title:     "Mr",
		FirstName: "Foo",
		LastName:  "FooFoo",
	})
	if err != nil {
		return nil, fmt.Errorf("could not marshal foo content: %v", err)
	}

	return &pkg.Message{
		ID:          "1234",
		MessageType: pkg.FooEventType,
		At:          time.Now(),
		Body:        data,
	}, nil
}

func createBarMessage() (*pkg.Message, error) {
	data, err := json.Marshal(pkg.BarContent{
		Address:  "404 BarNotFound St",
		PostCode: "9999",
		Region:   "FooBarLand",
	})
	if err != nil {
		return nil, fmt.Errorf("could not marshal bar content: %v", err)
	}
	return &pkg.Message{
		ID:          "1234",
		MessageType: pkg.BarEventType,
		At:          time.Now(),
		Body:        data,
	}, nil
}
