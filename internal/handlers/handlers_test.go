package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type postData struct {
	key   string
	value string
}

var TestCases = []struct {
	name               string
	url                string
	method             string
	expectedStatusCode int
	params             []postData
}{
	{"home", "/", "GET", http.StatusOK, []postData{}},
	{"singlebed", "/singlebed", "GET", http.StatusOK, []postData{}},
	{"coed", "/coed", "GET", http.StatusOK, []postData{}},
	{"highland", "/highland", "GET", http.StatusOK, []postData{}},
	{"reservation", "/reservation", "GET", http.StatusOK, []postData{}},
	{"contact", "/contact", "GET", http.StatusOK, []postData{}},
	{"make-reservation", "/make-reservation", "GET", http.StatusOK, []postData{}},
}

func TestHandlers(t *testing.T) {
	routes := getRoutes()
	testServer := httptest.NewTLSServer(routes)
	defer testServer.Close()

	// Table Testing
	for _, tc := range TestCases {
		if tc.method == "GET" {
			resp, err := testServer.Client().Get(testServer.URL + tc.url)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}
			if tc.expectedStatusCode != resp.StatusCode {
				t.Errorf("Excepted Code was %d, but got %d for %s", tc.expectedStatusCode, resp.StatusCode, tc.name)
			}
		}
	}

}
