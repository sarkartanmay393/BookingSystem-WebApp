package render

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/justinas/nosurf"
	"github.com/sarkartanmay393/RoomReservation-WebApp/internal/config"
	"github.com/sarkartanmay393/RoomReservation-WebApp/internal/models"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

var functions map[string]interface{}
var pathToTemplates = "./../../web/templates"

var appConf *config.AppConfig

// AttachConfig sets application config locally.
func AttachConfig(a *config.AppConfig) {
	appConf = a

	models.AttachConfigToTemplateData(a)
}

// AddDefaultData adds data that I want in every page of our web app.
func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.SuccessMessage = appConf.SessionManager.PopString(r.Context(), "success")
	td.ErrorMessage = appConf.SessionManager.PopString(r.Context(), "error")
	td.WarningMessage = appConf.SessionManager.PopString(r.Context(), "warning")
	td.CSRFToken = nosurf.Token(r)
	return td
}

// TemplateRender renders a specific template.
func TemplateRender(w http.ResponseWriter, r *http.Request, tmpl string, templateData *models.TemplateData) error {
	// Get template cache from application config that we have already got inside "appConfig" variable.
	// UseCache is used to recreate templateCache or setting existing template cache.
	var tc map[string]*template.Template
	var err error
	if appConf.UseCache {
		tc = appConf.TemplateCache
		//log.Println("Used existing template cache.")
	} else {
		tc, err = CreateTemplateCache()
		//appConf.UseCache = true
		//log.Println("Create new template cache.")
		if err != nil {
			log.Printf("Tried to CreateTemplateCache() and failed with err: %v", err)
			return err
		}
	}
	// Problem: There is a problem of how to use UseCache to only create new tc when there is changes in .tmpl files.
	t, ok := tc[tmpl] // Taking exact template as user requested.
	if !ok {
		return errors.New("not getting template from template cache")
	}
	templateData = AddDefaultData(templateData, r) // Adds default data.
	buffer := new(bytes.Buffer)
	_ = t.Execute(buffer, templateData)
	_, err = buffer.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template buffer to browser")
		return err
	}

	return nil
}

// CreateTemplateCache creates a map of web.
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}
	//pages, err := filepath.Glob(fmt.Sprintf("%s/*.page.tmpl", pathToTemplates))
	pages, err := filepath.Glob(fmt.Sprintf("%s/*.page.tmpl", pathToTemplates))
	if err != nil {
		return myCache, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		//fmt.Println(name)
		templateSet, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		layouts, err := filepath.Glob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates))
		if err != nil {
			return myCache, err
		}
		if len(layouts) > 0 {
			templateSet, err = templateSet.ParseGlob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates))
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = templateSet
	}
	return myCache, nil
}
