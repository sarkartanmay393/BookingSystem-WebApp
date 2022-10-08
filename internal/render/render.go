package render

import (
	"bytes"
	"fmt"
	"github.com/sarkartanmay393/RoomReservation-WebApp/internal/config"
	"github.com/sarkartanmay393/RoomReservation-WebApp/internal/models"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

var functions map[string]interface{}
var pathToTemplates = "../../web/templates"

var appConf *config.AppConfig

// AttachConfig sets application config locally.
func AttachConfig(a *config.AppConfig) {
	appConf = a

	models.AttachConfigToTemplateData(a)
}

// TemplateRender renders a specific template.
func TemplateRender(w http.ResponseWriter, r *http.Request, tmpl string, templateData *models.TemplateData) {
	// Get template cache from application config that we have already got inside "appConfig" variable.
	// UseCache is used to recreate templateCache or setting existing template cache.
	var tc map[string]*template.Template
	if appConf.UseCache {
		tc = appConf.TemplateCache
		//log.Println("Used existing template cache.")
	} else {
		tc, _ = CreateTemplateCache()
		//appConf.UseCache = true
		//log.Println("Create new template cache.")
	}
	// Problem: There is a problem of how to use UseCache to only create new tc when there is changes in .tmpl files.
	t, ok := tc[tmpl] // Taking exact template as user requested.
	if !ok {
		log.Fatalf("Error in TemplateRender() function.\n")
	}
	templateData = templateData.AddDefaultData(r) // Adds default data.
	buffer := new(bytes.Buffer)
	_ = t.Execute(buffer, templateData)
	_, err := buffer.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template buffer to browser")
	}
}

// CreateTemplateCache creates a map of web.
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}
	pages, err := filepath.Glob(fmt.Sprintf("%s/*.page.tmpl", pathToTemplates))
	if err != nil {
		return myCache, err
	}
	for _, page := range pages {

		name := filepath.Base(page) // "/template/home.page.tmpl" -> "home/page.tmpl"
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
