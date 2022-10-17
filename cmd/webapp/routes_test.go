package main

import (
	"github.com/go-chi/chi"
	"github.com/sarkartanmay393/RoomReservation-WebApp/internal/config"
	"testing"
)

func TestRouter(t *testing.T) {
	var testConfig config.AppConfig
	mux := router(&testConfig)
	switch v := mux.(type) {
	case *chi.Mux:
	// do nothing
	default:
		t.Errorf("Failure in testing routes function.")
		panic(v)
	}
}
