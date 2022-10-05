package handlers

import (
	"encoding/json"
	"github.com/justinas/nosurf"
	"github.com/sarkartanmay393/RoomReservation-WebApp/internal/config"
	"github.com/sarkartanmay393/RoomReservation-WebApp/internal/form"
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

	repo.app.SessionManager.Put(r.Context(), "chosenDates", &models.ChosenDates{Start: start, End: end})
	http.Redirect(w, r, "/make-reservation", http.StatusSeeOther)

	//w.Write([]byte(fmt.Sprintf("Start date is %s and End date is %s", start, end)))
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

// MakeReservationHandler handles form page with get and post request.
func (repo *Repository) MakeReservationHandler(w http.ResponseWriter, r *http.Request) {
	var emptyReservation models.Reservation
	data := make(map[string]interface{})
	data["reservation"] = emptyReservation

	render.TemplateRender(w, r, "make-reservation.page.tmpl", &models.TemplateData{
		Form: form.New(nil),
		Data: data,
	})
}

// PostMakeReservationHandler handles form page with get and post request.
func (repo *Repository) PostMakeReservationHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}

	reservation := &models.Reservation{
		FirstName: r.Form.Get("first_name"),
		LastName:  r.Form.Get("last_name"),
		Email:     r.Form.Get("email"),
		Phone:     r.Form.Get("phone"),
	}

	form := form.New(r.PostForm)
	//form.Has("first_name", r)
	form.Required("first_name", "last_name", "email", "phone")

	form.MinLength("first_name", 3, r)
	form.MinLength("last_name", 3, r)
	form.MinLength("phone", 10, r)
	form.IsEmail("email")

	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation

		render.TemplateRender(w, r, "make-reservation.page.tmpl", &models.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}
	repo.app.SessionManager.Put(r.Context(), "reservation", reservation)
	http.Redirect(w, r, "/reservation-summary", http.StatusSeeOther)
}

func (repo *Repository) ReservationSummaryHandler(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	reservation, ok := repo.app.SessionManager.Get(r.Context(), "reservation").(*models.Reservation)
	chosenDates, ok1 := repo.app.SessionManager.Get(r.Context(), "chosenDates").(*models.ChosenDates)
	if ok == ok1 && ok == true {
		data["reservation"] = reservation
		data["chosenDates"] = chosenDates
		render.TemplateRender(w, r, "reservation-summary.page.tmpl", &models.TemplateData{
			Data: data,
		})
		return
	}
	log.Println("Failed to retrieve reservation data")
	data["reservation"] = &models.Reservation{}
	data["chosenDates"] = &models.ChosenDates{}
	render.TemplateRender(w, r, "reservation-summary.page.tmpl", &models.TemplateData{
		Data: data,
	})

}
