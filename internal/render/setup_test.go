package render

import (
	"encoding/gob"
	"github.com/alexedwards/scs/v2"
	"github.com/sarkartanmay393/RoomReservation-WebApp/internal/config"
	"github.com/sarkartanmay393/RoomReservation-WebApp/internal/models"
	"log"
	"net/http"
	"os"
	"testing"
	"time"
)

var appConfTest config.AppConfig
var sessionTest *scs.SessionManager

func TestMain(m *testing.M) {
	gob.Register(&models.Reservation{})
	gob.Register(&models.ChosenDates{})

	// Starting of Logging information.
	appConfTest.InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	appConfTest.ErrorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Session Management Implementation
	sessionTest = scs.New()
	sessionTest.Lifetime = 24 * time.Hour
	sessionTest.Cookie.Persist = true
	sessionTest.Cookie.SameSite = http.SameSiteLaxMode
	sessionTest.Cookie.Secure = false
	appConfTest.SessionManager = sessionTest // Transfers this session object to app config.
	appConf = &appConfTest

	os.Exit(m.Run())
}

// Custom Object of http.ResponseWriter
type httpResponseWriter struct {
}

func (rw *httpResponseWriter) Header() http.Header {
	var h http.Header
	return h
}
func (rw *httpResponseWriter) Write(b []byte) (int, error) {
	length := len(b)
	return length, nil
}
func (rw *httpResponseWriter) WriteHeader(statusCode int) {
}
