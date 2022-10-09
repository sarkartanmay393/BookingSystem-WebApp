package models

import (
	"net/http"
	"testing"
)

func TestTemplateData_AddDefaultData(t *testing.T) {
	var td TemplateData
	request, err := getSessionActivated()
	if err != nil {
		t.Errorf("Failed to activate session at line:10, Error: %v", err)
	}
	sessionTest.Put(request.Context(), "success", "Success")
	sessionTest.Put(request.Context(), "error", "Error")
	sessionTest.Put(request.Context(), "warning", "Warning")
	tData := td.AddDefaultData(request)

	if tData.SuccessMessage != "Success" {
		t.Errorf("Failed to get Success Message correctly, Msg: %v", tData.SuccessMessage)
	}
	if tData.ErrorMessage != "Error" {
		t.Errorf("Failed to get Error Message correctly, Msg: %v", tData.ErrorMessage)
	}
	if tData.WarningMessage != "Warning" {
		t.Errorf("Failed to get Warning Message correctly, Msg: %v", tData.WarningMessage)
	}
}

func getSessionActivated() (*http.Request, error) {
	r, err := http.NewRequest("GET", "/testRequest", nil)
	if err != nil {
		return nil, err
	}
	ctx := r.Context()
	ctx, err = sessionTest.Load(ctx, r.Header.Get("X-Session"))
	if err != nil {
		return nil, err
	}
	r = r.WithContext(ctx)

	return r, nil
}
