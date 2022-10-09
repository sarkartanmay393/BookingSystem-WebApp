package form

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"net/url"
	"strings"
)

// Form is a custom structure for our use case with url.Values object.
type Form struct {
	url.Values
	Errors errors
}

// New returns newly initialised form address.
func New(data url.Values) *Form {
	return &Form{data, errors(map[string][]string{})}
}

// Has returns if there is the specific value is in the form or not.
func (f *Form) Has(field string) bool {
	exist := f.Get(field)
	if len(exist) == 0 {
		f.Errors.Add(field, "This field can not be blank")
		return false
	}
	return true
}

// Valid returns true if there is no error.
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

// Required adds error if there is blank field.
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field can not be blank")
		}
	}
}

func (f *Form) MinLength(field string, length int) bool {
	value := f.Get(field)
	if len(value) < length {
		f.Errors.Add(field, fmt.Sprintf("This length must be atleast %d", length))
		return false
	}
	return true
}

func (f *Form) IsEmail(field string) bool {
	if !govalidator.IsEmail(f.Get(field)) {
		f.Errors.Add(field, "You have entered a invalid email")
		return false
	}
	return true
}
