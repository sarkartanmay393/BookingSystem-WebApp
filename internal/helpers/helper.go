package helpers

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/sarkartanmay393/RoomReservation-WebApp/internal/config"
)

var app *config.AppConfig

func ConnectToHelpers(a *config.AppConfig) {
	app = a
}

func ClientError(w http.ResponseWriter, statusCode int) {
	app.InfoLog.Println("Client error with status code: ", statusCode)
	http.Error(w, http.StatusText(statusCode), statusCode)
}

func ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.ErrorLog.Println(trace)
	http.Error(w, fmt.Sprintf("%s\tERR: %s", http.StatusText(http.StatusInternalServerError), err), http.StatusInternalServerError)
}

func StatusText(i int) {
	panic("unimplemented")
}
