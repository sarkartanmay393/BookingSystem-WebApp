package models

import (
	"github.com/justinas/nosurf"
	"github.com/sarkartanmay393/RoomReservation-WebApp/internal/config"
	"github.com/sarkartanmay393/RoomReservation-WebApp/internal/form"
	"net/http"
)

// TemplateData is to be sent from hanlders to web.
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int

	Token     string
	CSRFToken string

	Form *form.Form
	Data map[string]interface{}

	SuccessMessage string
	ErrorMessage   string
	WarningMessage string
}

var appConfig *config.AppConfig

func AttachConfigToTemplateData(appCon *config.AppConfig) {
	appConfig = appCon
}

// AddDefaultData adds data that I want in every page of our web app.
func (td *TemplateData) AddDefaultData(r *http.Request) *TemplateData {
	td.SuccessMessage = appConfig.SessionManager.PopString(r.Context(), "success")
	td.ErrorMessage = appConfig.SessionManager.PopString(r.Context(), "error")
	td.WarningMessage = appConfig.SessionManager.PopString(r.Context(), "warning")
	td.CSRFToken = nosurf.Token(r)
	return td
}
