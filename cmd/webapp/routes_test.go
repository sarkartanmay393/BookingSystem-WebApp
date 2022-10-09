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
		t.Errorf("Type is not *chi.Mux in TestRouter() line:16 : %T", v)
	}
}
