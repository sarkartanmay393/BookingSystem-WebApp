package render

import (
	"github.com/sarkartanmay393/RoomReservation-WebApp/internal/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddDefaultData(t *testing.T) {
	var td *models.TemplateData
	request, err := getSessionActivated()
	if err != nil {
		t.Errorf("Failed to activate session at line:10, Error: %v", err)
	}
	sessionTest.Put(request.Context(), "success", "Success")
	sessionTest.Put(request.Context(), "error", "Error")
	sessionTest.Put(request.Context(), "warning", "Warning")
	td = AddDefaultData(td, request)

	if td.SuccessMessage != "Success" {
		t.Errorf("Failed to get Success Message correctly, Msg: %v", td.SuccessMessage)
	}
	if td.ErrorMessage != "Error" {
		t.Errorf("Failed to get Error Message correctly, Msg: %v", td.ErrorMessage)
	}
	if td.WarningMessage != "Warning" {
		t.Errorf("Failed to get Warning Message correctly, Msg: %v", td.WarningMessage)
	}
}

//func TestTemplateRender(t *testing.T) {
//	rTest, err := getSessionActivated()
//	var wTest httpResponseWriter
//	if err != nil {
//		t.Error(err)
//	}
//	cache, err := CreateTemplateCache()
//	if err != nil {
//		t.Error(err)
//	}
//	appConfTest.TemplateCache = cache
//
//	err = TemplateRender(&wTest, rTest, "coed.page.tmpl", &models.TemplateData{})
//	if err != nil {
//		t.Errorf("Error: %v", err)
//	}
//}

//func TestCreateTemplateCache(t *testing.T) {
//	//pathToTemplates = "./../../web/templates"
//	_, err := CreateTemplateCache()
//	if err != nil {
//		t.Errorf("Failed to create templates caches, Error: %v", err)
//	}
//}

func getSessionActivated() (*http.Request, error) {
	r := httptest.NewRequest("GET", "/", nil)
	ctx := r.Context()
	ctx, err := appConfTest.SessionManager.Load(ctx, r.Header.Get("X-Session"))
	if err != nil {
		return nil, err
	}
	r = r.WithContext(ctx)
	return r, nil
}
