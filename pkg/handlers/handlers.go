package handlers

import (
	"github.com/sarkartanmay393/BookingSystem-WebApp/pkg/config"
	"github.com/sarkartanmay393/BookingSystem-WebApp/pkg/models"
	"github.com/sarkartanmay393/BookingSystem-WebApp/pkg/render"
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
func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

// HomeHandler handles main page on "/".
func (repo *Repository) HomeHandler(w http.ResponseWriter, r *http.Request) {
	sMap := make(map[string]string) // String map to be passed.
	sMap["Test"] = "Hello WORLD!"

	remoteIP := r.RemoteAddr                                        // Getting IP address from request body.
	repo.app.SessionManager.Put(r.Context(), "remote_ip", remoteIP) // Saving IP on session manager.

	render.TemplateRender(w, "home.page.tmpl", &models.TemplateData{
		StringMap: sMap,
	})
}

// FormHandler handles main page on "/form".
func (repo *Repository) FormHandler(w http.ResponseWriter, r *http.Request) {
	remoteIP := repo.app.SessionManager.GetString(r.Context(), "remote_ip") // Parsing IP from session manager.
	sMap := make(map[string]string)                                         // String map to be passed.
	sMap["remote_ip"] = remoteIP

	render.TemplateRender(w, "form.page.tmpl", &models.TemplateData{
		StringMap: sMap,
	})
}
