package main

import (
	"net/http"
	"testing"
)

func TestCSRFCheck(t *testing.T) {
	var demoHandler myHandler
	h := CSRFCheck(&demoHandler)

	switch v := h.(type) {
	case http.Handler:
	// do nothing
	default:
		t.Errorf("type is not http.Handler in TestCSRFCheck() line:16 : %T", v)
	}
}
