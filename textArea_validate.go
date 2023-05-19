package input

import "regexp"

// validación con datos de entrada
func (t textArea) ValidateField(data_in string, skip_validation bool) bool {
	if !skip_validation {

		pvalid := regexp.MustCompile(t.pattern)

		return pvalid.MatchString(data_in)

	} else {
		return true
	}
}
