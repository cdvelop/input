package input

import (
	"strconv"
	"strings"
)

// validación con datos de entrada
func (r rut) ValidateField(data_in string, skip_validation bool, options ...string) (err string) {
	if !skip_validation {

		for _, doc := range options {
			if doc == "ex" {
				return r.dni.Validate(data_in)
			} else {
				return r.runValidate(data_in)
			}
		}

		if r.dni_mode {
			if !strings.Contains(data_in, `-`) {
				return r.dni.Validate(data_in)
			}
		}

		return r.runValidate(data_in)

	}

	return ""
}

// RUT validate formato "7863697-1"
func (r rut) runValidate(rin string) (err string) {
	data, onlyRun, err := RunData(rin)
	if err != "" {
		return err
	}
	// log.Printf("DATA: [%v] RUN:[%v] TAMAÑO DATA: [%v]\n", data, onlyRun, len(data))

	if data[0][0:1] == "0" {
		return "primer dígito no puede ser 0"
	}

	dv := DvRut(onlyRun)

	// log.Printf("DÍGITO: %v DV DATA: %v\n", dv, strings.ToLower(data[1]))

	if dv != strings.ToLower(data[1]) {
		return "dígito verificador " + data[1] + " inválido"

	}

	return ""
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

func RunData(runIn string) (data []string, onlyRun int, err string) {

	if runIn == "" || runIn == " " {
		return nil, 0, "rut sin información"
	}

	if !strings.Contains(runIn, "-") {
		return nil, 0, "rut incorrecto"
	}

	data = strings.Split(string(runIn), "-")
	// fmt.Println("TAMAÑO", len(data), "RUT DATA -:", data)
	var e error
	onlyRun, e = strconv.Atoi(data[0])
	if e != nil {
		err = "caracteres no permitidos"
	}

	return
}
