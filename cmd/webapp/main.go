package main

import (
	"encoding/gob"
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/sarkartanmay393/RoomReservation-WebApp/internal/config"
	"github.com/sarkartanmay393/RoomReservation-WebApp/internal/handlers"
	"github.com/sarkartanmay393/RoomReservation-WebApp/internal/models"
	"github.com/sarkartanmay393/RoomReservation-WebApp/internal/render"
	"log"
	"net/http"
	"os"
	"time"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	err := RunMain()
	if err != nil {
		log.Fatalln("Failed to execute runMain() function in main.go file.")
	}

	log.Println("Server started on port 8080 💫")
	// Serving and Handling web with help of pat pkg.
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%v", os.Getenv("PORT")),
		Handler: router(&app),
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatalln("Listen and serving error occurred")
	}
}

func RunMain() error {
	gob.Register(&models.Reservation{})
	gob.Register(&models.ChosenDates{})
	// Creating template cache for the whole app to get started.
	var err error
	app.TemplateCache, err = render.CreateTemplateCache()
	app.UseCache = true      // Manual value setup in app config.
	app.InProduction = false // Manual value setup in app config.
	if err != nil {
		log.Print("Error creating template cache\n") // No newline because of testing.
		return err
	}
	render.AttachConfig(&app)                     // appConfig is transferred to render.go file.
	temporaryRepo := handlers.CreateNewRepo(&app) // Creates a new Repo with global appConfig to be transferred.
	handlers.AttachRepo(temporaryRepo)            // appConfig is transferred to handlers.go file.
	// Now application config will be available in render.go and handlers.go file.

	// Session Management Implementation
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.SessionManager = session // Transfers this session object to app config.

	return nil
}
