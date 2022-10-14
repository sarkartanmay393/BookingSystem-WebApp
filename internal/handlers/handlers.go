package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/sarkartanmay393/RoomReservation-WebApp/internal/config"
	"github.com/sarkartanmay393/RoomReservation-WebApp/internal/driver"
	"github.com/sarkartanmay393/RoomReservation-WebApp/internal/form"
	"github.com/sarkartanmay393/RoomReservation-WebApp/internal/helpers"
	"github.com/sarkartanmay393/RoomReservation-WebApp/internal/models"
	"github.com/sarkartanmay393/RoomReservation-WebApp/internal/render"
)

// Repo is local variable for global configuration.
var Repo *Repository

// Repository holds global application configurations.
type Repository struct {
	app *config.AppConfig
	db  *driver.DB
}

// CreateNewRepo return a new Repository with provided appConfig.
func CreateNewRepo(a *config.AppConfig, db *driver.DB) *Repository {
	return &Repository{
		app: a,
		db:  db,
	}
}

// AttachRepo attaches Repository inside handler.go file.
func AttachRepo(r *Repository) {
	Repo = r
}

// HomeHandler handles main page on "/".
func (repo *Repository) HomeHandler(w http.ResponseWriter, r *http.Request) {
	render.TemplateRender(w, r, "home.page.tmpl", &models.TemplateData{})
}

// CoedHandler handles main page on "/coed".
func (repo *Repository) CoedHandler(w http.ResponseWriter, r *http.Request) {
	render.TemplateRender(w, r, "coed.page.tmpl", &models.TemplateData{})
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
	repo.app.SessionManager.Put(r.Context(), "sdates", r.Form.Get("start-date"))
	repo.app.SessionManager.Put(r.Context(), "edates", r.Form.Get("end-date"))

	start, err := time.Parse("02-01-2006", r.Form.Get("start-date"))
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	end, err := time.Parse("02-01-2006", r.Form.Get("end-date"))
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	rooms, err := repo.db.SearcAvailabilityByDates(start, end)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	// for _, r := range rooms {
	// 	log.Printf("ID: %v, ROOM: %v\n", r.ID, r.RoomName)
	// }

	if len(rooms) == 0 {
		repo.app.SessionManager.Put(r.Context(), "error", "No available rooms in those dates")
		http.Redirect(w, r, "/reservation", http.StatusSeeOther)
		return
	}

	repo.app.SessionManager.Put(r.Context(), "rooms", rooms)
	repo.app.SessionManager.Put(r.Context(), "chosenDates", &models.ChosenDates{Start: start, End: end})
	http.Redirect(w, r, "/make-reservation", http.StatusSeeOther)
	// w.Write([]byte(fmt.Sprintf("Start date is %s and End date is %s", start, end)))
}

// MakeReservationHandler handles form page with get and post request.
func (repo *Repository) MakeReservationHandler(w http.ResponseWriter, r *http.Request) {
	var emptyReservation models.Reservation
	data := make(map[string]interface{})
	data["reservation"] = emptyReservation

	var ok bool
	_, ok = repo.app.SessionManager.Get(r.Context(), "chosenDates").(*models.ChosenDates)
	if !ok {
		repo.app.ErrorLog.Println("not found chosen dates")
		repo.app.SessionManager.Put(r.Context(), "warning", "No chosen dates")
		http.Redirect(w, r, "/reservation", http.StatusTemporaryRedirect)
		return
	}

	data["sdates"] = repo.app.SessionManager.PopString(r.Context(), "sdates")
	data["edates"] = repo.app.SessionManager.PopString(r.Context(), "edates")
	data["rooms"] = repo.app.SessionManager.Pop(r.Context(), "rooms")

	render.TemplateRender(w, r, "make-reservation.page.tmpl", &models.TemplateData{
		Form: form.New(nil),
		Data: data,
	})
}

// PostMakeReservationHandler handles form page with get and post request.
func (repo *Repository) PostMakeReservationHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	dates, ok := repo.app.SessionManager.Get(r.Context(), "chosenDates").(*models.ChosenDates)
	if !ok {
		repo.app.SessionManager.Put(r.Context(), "warning", "Not found chosen dates")
	}

	form := form.New(r.PostForm)
	form.Required("first_name", "last_name", "email", "phone")
	form.MinLength("first_name", 3)
	form.MinLength("last_name", 3)
	form.MinLength("phone", 10)
	form.IsEmail("email")

	reservation := &models.Reservation{
		FirstName: r.Form.Get("first_name"),
		LastName:  r.Form.Get("last_name"),
		Email:     form.Get("email"),
		Phone:     form.Get("phone"),
		StartDate: dates.Start,
		EndDate:   dates.End,
		RoomID:    1,
	}

	hotels := r.Form["hotels"]
	fmt.Print(hotels)

	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation

		render.TemplateRender(w, r, "make-reservation.page.tmpl", &models.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}

	reservation.ID, err = repo.db.InsertReservation(reservation)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	roomrestriction := &models.RoomRestriction{
		StartDate:     dates.Start,
		EndDate:       dates.End,
		RoomID:        reservation.RoomID,
		ReservationID: reservation.ID,
		RestrictionID: 1,
	}
	err = repo.db.InsertRoomRestriction(roomrestriction)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	repo.app.SessionManager.Put(r.Context(), "reservation", reservation)
	http.Redirect(w, r, "/reservation-summary", http.StatusSeeOther)
}

func (repo *Repository) ReservationSummaryHandler(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	reservation, ok := repo.app.SessionManager.Pop(r.Context(), "reservation").(*models.Reservation)
	if ok {
		data["reservation"] = reservation
		render.TemplateRender(w, r, "reservation-summary.page.tmpl", &models.TemplateData{
			Data: data,
		})
		return
	}

	repo.app.ErrorLog.Println("Failed to retrieve reservation data from session")
	repo.app.SessionManager.Put(r.Context(), "error", "Not found user information")
	http.Redirect(w, r, "/reservation", http.StatusTemporaryRedirect)
}

// ContactHandler handles main page on "/contact".
func (repo *Repository) ContactHandler(w http.ResponseWriter, r *http.Request) {
	render.TemplateRender(w, r, "contact.page.tmpl", &models.TemplateData{})
}

// HighlandHandler handles main page on "/highland".
func (repo *Repository) HighlandHandler(w http.ResponseWriter, r *http.Request) {
	render.TemplateRender(w, r, "highland.page.tmpl", &models.TemplateData{})
}
