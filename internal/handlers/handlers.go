package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/justinas/nosurf"
	"github.com/sarkartanmay393/RoomReservation-WebApp/internal/config"
	"github.com/sarkartanmay393/RoomReservation-WebApp/internal/models"
	"github.com/sarkartanmay393/RoomReservation-WebApp/internal/render"
	"log"

	"net/http"
)

// Repo is local variable for global configuration.
var Repo *Repository

// Repository holds global application configurations.
type Repository struct {
	app *config.AppConfig
}

// CreateNewRepo return a new Repository with provided appConfig.
func CreateNewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		app: a,
	}
}

// AttachRepo attaches Repository inside handler.go file.
func AttachRepo(r *Repository) {
	Repo = r
}

// AddDefaultData adds data that I want in every page of our web app.
func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.CSRFToken = nosurf.Token(r)
	return td
}

// HomeHandler handles main page on "/".
func (repo *Repository) HomeHandler(w http.ResponseWriter, r *http.Request) {
	sMap := make(map[string]string) // String map to be passed.
	sMap["Test"] = "Hello WORLD!"

	remoteIP := r.RemoteAddr                                        // Getting IP address from request body.
	repo.app.SessionManager.Put(r.Context(), "remote_ip", remoteIP) // Saving IP on session manager.

	render.TemplateRender(w, r, "home.page.tmpl", &models.TemplateData{
		StringMap: sMap,
	})
}

// CoedHandler handles main page on "/coed".
func (repo *Repository) CoedHandler(w http.ResponseWriter, r *http.Request) {
	remoteIP := repo.app.SessionManager.GetString(r.Context(), "remote_ip") // Parsing IP from session manager.
	sMap := make(map[string]string)                                         // String map to be passed.
	sMap["remote_ip"] = remoteIP

	render.TemplateRender(w, r, "coed.page.tmpl", &models.TemplateData{
		StringMap: sMap,
	})
}

// SinglebedHandler handles main page on "/singlebed".
func (repo *Repository) SinglebedHandler(w http.ResponseWriter, r *http.Request) {
	render.TemplateRender(w, r, "singlebed.page.tmpl", &models.TemplateData{})
}

// ReservationHandler handles main page on "/reservation".
func (repo *Repository) ReservationHandler(w http.ResponseWriter, r *http.Request) {
	render.TemplateRender(w, r, "reservation.page.tmpl", &models.TemplateData{})
}

// PostReservationHandler handles main page on "/reservation".
func (repo *Repository) PostReservationHandler(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start-date")
	end := r.Form.Get("end-date")
	w.Write([]byte(fmt.Sprintf("Start date is %s and End date is %s", start, end)))
}

// Custom jsonResponse structure for our own custom responses.
type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

// AvailabilityHandler handles post request on "/reservation-json" and send json response.
func (repo *Repository) AvailabilityHandler(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		OK:      true,
		Message: "Available!",
	}

	out, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		log.Print(err)
	}
	//log.Println(string(out))

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

// ContactHandler handles main page on "/contact".
func (repo *Repository) ContactHandler(w http.ResponseWriter, r *http.Request) {
	render.TemplateRender(w, r, "contact.page.tmpl", &models.TemplateData{})
}

// HighlandHandler handles main page on "/highland".
func (repo *Repository) HighlandHandler(w http.ResponseWriter, r *http.Request) {
	render.TemplateRender(w, r, "highland.page.tmpl", &models.TemplateData{})
}
