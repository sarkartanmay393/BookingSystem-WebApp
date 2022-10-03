package form

type errors map[string][]string

// Add appends error message.
func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

// Get returns first error message.
func (e errors) Get(field string) string {
	err := e[field]
	if len(err) == 0 {
		return ""
	}
	return err[0]
}
