package models

import (
	"github.com/sarkartanmay393/RoomReservation-WebApp/internal/form"
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
