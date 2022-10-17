package driver

import (
	"encoding/gob"
	"github.com/sarkartanmay393/RoomReservation-WebApp/internal/models"
	"log"
	"os"
	"testing"
)

var db *DB

func TestMain(m *testing.M) {

	gob.Register(&models.User{})
	gob.Register([]models.Room{})
	gob.Register(&models.Reservation{})
	gob.Register(&models.Restriction{})
	gob.Register(&models.RoomRestriction{})

	log.Println("Connecting to database...")

	dsn := "host=localhost port=5432 dbname=roomreservation user=postgres password=Tanmay3597!"
	db, err := ConnectSQL(dsn)
	if err != nil {
		log.Fatalln("unable to connect database: ", err)
	}
	log.Println("Connected to database!")

	defer db.SQL.Close()
	os.Exit(m.Run())
}
