package config

import (
	"github.com/alexedwards/scs/v2"
	"github.com/sarkartanmay393/RoomReservation-WebApp/internal/models"
	"log"
	"text/template"
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
}
