package input

import (
	"regexp"
	"strconv"
	"strings"
)

// validación con datos de entrada
func (r rut) ValidateField(data_in string, skip_validation bool, options ...string) bool {
	if !skip_validation {

		for _, doc := range options {

			if doc == "ex" {
				pvalid := regexp.MustCompile(r.Pattern)
				return pvalid.MatchString(data_in)
			}
		}

		if r.dni_mode {
			if !strings.Contains(data_in, `-`) {
				pvalid := regexp.MustCompile(r.Pattern)
				return pvalid.MatchString(data_in)
			}
		}

		return r.runValidate(data_in)

	} else {
		return true
	}
}

// RUT validate formato "7863697-1"
func (r rut) runValidate(rin string) bool {
	data, onlyRun, ok := runData(rin)
	if !ok {
		return false
	}
	// log.Printf("DATA: [%v] RUN:[%v] TAMAÑO DATA: [%v]\n", data, onlyRun, len(data))

	if len(data) <= 1 {
		return false
	}

	if data[0] == "" {
		return false
	}

	if data[0][0:1] == "0" {
		return false
	}

	dv := DvRut(onlyRun)

	// log.Printf("DÍGITO: %v DV DATA: %v\n", dv, strings.ToLower(data[1]))

	return dv == strings.ToLower(data[1])
}

// DvRut retorna dígito verificador de un run
func DvRut(rut int) string {
	var sum = 0
	var factor = 2
	for ; rut != 0; rut /= 10 {
		sum += rut % 10 * factor
		if factor == 7 {
			factor = 2
		} else {
			factor++
		}
	}

	if val := 11 - (sum % 11); val == 11 {
		return "0"
	} else if val == 10 {
		return "k"
	} else {
		return strconv.Itoa(val)
	}
}
