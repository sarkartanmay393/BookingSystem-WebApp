package models

import (
	"github.com/justinas/nosurf"
	"net/http"
)

// TemplateData is to be sent from hanlders to web.
type TemplateData struct {
	StringMap      map[string]string
	IntMap         map[string]int
	Token          string
	SuccessMessage string
	FailMessage    string
	CSRFToken      string
}

// AddDefaultData adds data that I want in every page of our web app.
func (td *TemplateData) AddDefaultData(r *http.Request) *TemplateData {
	td.CSRFToken = nosurf.Token(r)
	return td
}
