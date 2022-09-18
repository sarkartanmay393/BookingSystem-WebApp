package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/sarkartanmay393/RoomReservation-WebApp/pkg/config"
	"github.com/sarkartanmay393/RoomReservation-WebApp/pkg/handlers"
	"net/http"
)

func router(a *config.AppConfig) http.Handler {
	mux := chi.NewRouter() // Instance of router using chi package.

	mux.Use(middleware.Recoverer)           // Tackle panic attack as a middleware.
	mux.Use(WriteToConsole)                 // Write new page load as a middleware.
	mux.Use(CSRFCheck)                      // Checks for Cross-site request forgery attacks.
	mux.Use(app.SessionManager.LoadAndSave) // Loads and saves session data.

	mux.Get("/", http.HandlerFunc(handlers.Repo.HomeHandler))     // Serve root page request.
	mux.Get("/form", http.HandlerFunc(handlers.Repo.FormHandler)) // Serve /form page request.

	fileServer := http.FileServer(http.Dir("./static/"))             // fileServer handles file system contents.
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer)) // Handles all files.

	return mux // Returns the http.handler for further use in main.go
}
