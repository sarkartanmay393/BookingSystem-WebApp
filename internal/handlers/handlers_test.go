package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
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
	{"reservation-summary", "/reservation-summary", "GET", http.StatusOK, []postData{}},

	{"reservation", "/reservation", "POST", http.StatusOK, []postData{
		{key: "start-date", value: "01-01-2022"},
		{key: "end-date", value: "01-01-2022"},
	}},
	{"reservation-json", "/reservation-json", "POST", http.StatusOK, []postData{
		{key: "start-date", value: "01-01-2022"},
		{key: "end-date", value: "01-01-2022"},
	}},
	{"make-reservation", "/make-reservation", "POST", http.StatusOK, []postData{
		{key: "first_name", value: "Tanmay"},
		{key: "last_name", value: "Sarkar"},
		{key: "email", value: "sarkartanmay393@gmail.com"},
		{key: "phone", value: "7609001122"},
	}},
}

func TestHandlers(t *testing.T) {
	routes := getRoutes()
	testServer := httptest.NewTLSServer(routes)
	defer testServer.Close() // It executes when whole function is done.

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
		} else {
			// For POST requests we have to turn off CSRF Check middleware.
			values := url.Values{}
			for _, params := range tc.params {
				values.Add(params.key, params.value)
			}
			resp, err := testServer.Client().PostForm(testServer.URL+tc.url, values)
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
