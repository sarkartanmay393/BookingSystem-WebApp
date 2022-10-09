package render

import (
	"github.com/sarkartanmay393/RoomReservation-WebApp/internal/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddDefaultData(t *testing.T) {
	var td = &models.TemplateData{
		StringMap:      nil,
		IntMap:         nil,
		Token:          "",
		CSRFToken:      "",
		Form:           nil,
		Data:           nil,
		SuccessMessage: "",
		ErrorMessage:   "",
		WarningMessage: "",
	}

	request, err := getSessionActivated()
	if err != nil {
		t.Errorf("Failed to activate session, Error: %v", err)
	}

	sessionTest.Put(request.Context(), "success", "Success")
	sessionTest.Put(request.Context(), "error", "Error")
	sessionTest.Put(request.Context(), "warning", "Warning")

	td = AddDefaultData(td, request)

	if td.SuccessMessage != "Success" {
		t.Errorf("Failed to get Success Message correctly, Msg: %s", td.SuccessMessage)
	}
	if td.ErrorMessage != "Error" {
		t.Errorf("Failed to get Error Message correctly, Msg: %s", td.ErrorMessage)
	}
	if td.WarningMessage != "Warning" {
		t.Errorf("Failed to get Warning Message correctly, Msg: %s", td.WarningMessage)
	}
}

func TestTemplateRender(t *testing.T) {
	rTest, err := getSessionActivated()
	var wTest httpResponseWriter

	if err != nil {
		t.Error(err)
	}
	cache, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)
	}
	appConfTest.TemplateCache = cache

	err = TemplateRender(&wTest, rTest, "coed.page.tmpl", &models.TemplateData{})
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}

func TestCreateTemplateCache(t *testing.T) {
	pathToTemplates = "./../../web/templates"
	_, err := CreateTemplateCache()
	if err != nil {
		t.Errorf("Failed to create templates caches, Error: %v", err)
	}
}

func TestAttachConfig(t *testing.T) {
	AttachConfig(&appConfTest)
}

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
