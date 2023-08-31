package input

import (
	"strconv"
	"strings"

	"github.com/cdvelop/model"
)

// validación con datos de entrada
func (r rut) ValidateField(data_in string, skip_validation bool, options ...string) error {
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

	return nil
}

// RUT validate formato "7863697-1"
func (r rut) runValidate(rin string) error {
	data, onlyRun, err := RunData(rin)
	if err != nil {
		return err
	}
	// log.Printf("DATA: [%v] RUN:[%v] TAMAÑO DATA: [%v]\n", data, onlyRun, len(data))

	if data[0][0:1] == "0" {
		return model.Error("primer dígito no puede ser 0")
	}

	dv := DvRut(onlyRun)

	// log.Printf("DÍGITO: %v DV DATA: %v\n", dv, strings.ToLower(data[1]))

	if dv != strings.ToLower(data[1]) {
		return model.Error("dígito verificador", data[1], "inválido")

	}

	return nil
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

func RunData(runIn string) (data []string, onlyRun int, err error) {

	if runIn == "" || runIn == " " {
		return nil, 0, model.Error("rut sin información")
	}

	if !strings.Contains(runIn, "-") {
		return nil, 0, model.Error("rut incorrecto")
	}

	data = strings.Split(string(runIn), "-")
	// fmt.Println("TAMAÑO", len(data), "RUT DATA -:", data)

	onlyRun, err = strconv.Atoi(data[0])
	if err != nil {
		err = model.Error("caracteres no permitidos")
	}

	return
}
