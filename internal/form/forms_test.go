package form

import (
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestNew(t *testing.T) {
	_ = New(url.Values{})
}

func TestForm_Has(t *testing.T) {
	// Test with invalid value.
	testForm := New(url.Values{})
	has := testForm.Has("random")
	if has {
		t.Errorf("It doesn't consist the value, hence error.")
	}

	// Test with valid value.
	values := url.Values{}
	values.Add("test", "success")
	testForm = New(values)
	has = testForm.Has("test")
	if !has {
		t.Errorf("It should consist the value, hence Error.")
	}
}

func TestForm_Valid(t *testing.T) {
	testForm := New(url.Values{})
	ok := testForm.Valid()
	if !ok {
		t.Errorf("Form Valid Test Failed")
	}
}

func TestForm_Required(t *testing.T) {
	values := url.Values{}
	values.Add("first_name", "Tanmay")
	values.Add("last_name", "Sarkar")
	values.Add("email", "sarkartanmay393@gmail.com")
	testForm := New(values)
	testForm.Required("first_name", "last_name", "email")
	if len(testForm.Errors) != 0 {
		t.Errorf("Form Required Test Failed with Valid Values")
	}

	testForm = New(url.Values{})
	testForm.Required("random")
	if len(testForm.Errors) == 0 {
		t.Errorf("Form Required Test Failed with Invalid Values")
	}

}

func TestForm_MinLength(t *testing.T) {
	r := httptest.NewRequest("GET", "/", nil)
	values := url.Values{}
	values.Add("first_name", "Tanmay")
	values.Add("last_name", "Sar")
	values.Add("phone", "123456789")
	values.Add("something", "")
	testForm := New(values)
	//ok := testForm.MinLength("first_name", 3, r)
	//if !ok {
	//	t.Errorf("Failed with valid values, Needed: 3 Provided: 6")
	//}
	//ok = testForm.MinLength("last_name", 3, r)
	//if !ok {
	//	t.Errorf("Failed with valid values, Needed: 3 Provided: 3")
	//}
	ok := testForm.MinLength("phone", 10, r)
	if ok {
		t.Errorf("Success with invalid values, Needed: 10 Provided: 9")
	}
	ok = testForm.MinLength("something", 3, r)
	if ok {
		t.Errorf("Failed with invalid values, Needed: 3 Provided: 0")
	}

}

func TestForm_IsEmail(t *testing.T) {
	values := url.Values{}
	values.Add("email", "sarkartanmay393@gmail.com")
	values.Add("email2", "sarkartanmay393gmail.com")
	testForm := New(values)
	ok := testForm.IsEmail("email")
	if !ok {
		t.Errorf("Failed with valid value")
	}
	ok = testForm.IsEmail("email2")
	if ok {
		t.Errorf("Failed with invalid value")
	}
}
