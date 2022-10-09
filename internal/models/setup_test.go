package models

import (
	"encoding/gob"
	"github.com/alexedwards/scs/v2"
	"github.com/sarkartanmay393/RoomReservation-WebApp/internal/config"
	"net/http"
	"os"
	"testing"
	"time"
)

var appConfigTest config.AppConfig
var sessionTest *scs.SessionManager

func TestMain(m *testing.M) {
	gob.Register(&Reservation{})
	gob.Register(&ChosenDates{})
	// Session Management Implementation
	sessionTest = scs.New()
	sessionTest.Lifetime = 24 * time.Hour
	sessionTest.Cookie.Persist = true
	sessionTest.Cookie.SameSite = http.SameSiteLaxMode
	sessionTest.Cookie.Secure = false

	appConfigTest.SessionManager = sessionTest // Transfers this session object to app config.
	appConfig = &appConfigTest

	os.Exit(m.Run())
}
