package input

import (
	"regexp"
)

// validación con datos de entrada
func (dayWord) ValidateField(data_in string, skip_validation bool) bool { //en realidad es YYYY-MM-DD
	if !skip_validation {

		pvalid := regexp.MustCompile(patternMonthDay)

		return pvalid.MatchString(data_in)

	} else {
		return true
	}
}
