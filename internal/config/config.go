package config

import (
	"log"
	"text/template"

	"github.com/alexedwards/scs/v2"
	"github.com/sarkartanmay393/RoomReservation-WebApp/internal/models"
)

type AppConfig struct {
	UseCache       bool
	TemplateCache  map[string]*template.Template
	InfoLog        *log.Logger
	ErrorLog       *log.Logger
	SessionManager *scs.SessionManager
	InProduction   bool
	MailChannel    chan models.MailData
	IsLogin        bool
	RoomLoaded     bool
}
