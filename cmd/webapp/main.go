package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/sarkartanmay393/RoomReservation-WebApp/internal/driver"

	"github.com/alexedwards/scs/v2"
	"github.com/sarkartanmay393/RoomReservation-WebApp/internal/config"
	"github.com/sarkartanmay393/RoomReservation-WebApp/internal/handlers"
	"github.com/sarkartanmay393/RoomReservation-WebApp/internal/helpers"
	"github.com/sarkartanmay393/RoomReservation-WebApp/internal/models"
	"github.com/sarkartanmay393/RoomReservation-WebApp/internal/render"
)

var portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	db, err := RunMain()
	if err != nil {
		log.Fatalln("Failed to execute runMain() function in main.go file.")
	}
	defer db.SQL.Close()

	os.Setenv("PORT", "8080")
	portNumber = fmt.Sprintf(":%v", os.Getenv("PORT"))

	log.Println("Server started on port 8080 💫")
	// Serving and Handling web with help of pat pkg.
	srv := &http.Server{
		Addr:    portNumber,
		Handler: router(&app),
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatalln("Listen and serving error occurred")
	}

}

func RunMain() (*driver.DB, error) {

	gob.Register(&models.Reservation{})
	gob.Register(&models.ChosenDates{})
	gob.Register(&models.Restriction{})
	gob.Register([]models.Room{})
	gob.Register(&models.User{})
	gob.Register(&models.RoomRestriction{})

	log.Println("Connecting to database...")
	dsn := "host=localhost port=5432 dbname=roomreservation user=postgres password=Tanmay3597!"
	db, err := driver.ConnectSQL(dsn)
	if err != nil {
		log.Fatalln("unable to connect database: ", err)
	}
	log.Println("Connected to database!")

	// Starting of Logging information.
	app.InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.ErrorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Creating template cache for the whole app to get started.
	app.TemplateCache, err = render.CreateTemplateCache()
	app.UseCache = true      // Manual value setup in app config.
	app.InProduction = false // Manual value setup in app config.
	if err != nil {
		log.Print("Error creating template cache\n") // No newline because of testing.
		return db, err
	}
	render.AttachConfig(&app)                         // appConfig is transferred to render.go file.
	temporaryRepo := handlers.CreateNewRepo(&app, db) // Creates Repo with appConfig and db connection pool to be transferred.
	handlers.AttachRepo(temporaryRepo)                // appConfig is transferred to handlers.go file.
	helpers.ConnectToHelpers(&app)                    // App config now to helpers package.
	// Now application config will be available in render.go and handlers.go file.

	// Session Management Implementation
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.SessionManager = session // Transfers this session object to app config.

	return db, nil
}
