package models

// TemplateData is to be sent from hanlders to web.
type TemplateData struct {
	StringMap      map[string]string
	IntMap         map[string]int
	Token          string
	SuccessMessage string
	FailMessage    string
	CSRFToken      string
}
