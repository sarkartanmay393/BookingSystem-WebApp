package config

import (
	"github.com/alexedwards/scs/v2"
	"log"
	"text/template"
)

type AppConfig struct {
	UseCache       bool
	TemplateCache  map[string]*template.Template
	LogFile        *log.Logger
	SessionManager *scs.SessionManager
	InProduction   bool
}
