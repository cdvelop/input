package input

import (
	"strconv"
)

// validación con datos de entrada
func (r rut) ValidateField(data_in string, skip_validation bool, options ...string) (err string) {
	if skip_validation {
		return ""
	}

	const hidden_err = "campo invalido"

	for _, doc := range options {
		if doc == "ex" {
			err = r.dni.Validate(data_in)
			if err != "" && r.hide_typing {
				return hidden_err
			}
			return
		} else {
			err = r.runValidate(data_in)
			if err != "" && r.hide_typing {
				return hidden_err
			}
			return
		}
	}

	if r.dni_mode {
		if String().Contains(data_in, `-`) == 0 {
			err = r.dni.Validate(data_in)
			if err != "" && r.hide_typing {
				return hidden_err
			}
			return
		}
	}

	err = r.runValidate(data_in)
	if err != "" && r.hide_typing {
		return hidden_err
	}

	return

}

const errCeroRut = "primer dígito no puede ser 0"

// RUT validate formato "7863697-1"
func (r rut) runValidate(rin string) (err string) {
	data, onlyRun, err := RunData(rin)
	if err != "" {
		return err
	}
	// log.Printf("DATA: [%v] RUN:[%v] TAMAÑO DATA: [%v]\n", data, onlyRun, len(data))

	if data[0][0:1] == "0" {
		return errCeroRut
	}

	dv := DvRut(onlyRun)

	// log.Printf("DÍGITO: %v DV DATA: %v\n", dv, strings.ToLower(data[1]))

	if dv != String().ToLowerCase(data[1]) {
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

const errRut01 = "datos ingresados insuficientes"
const errGuionRut = "guion (-) dígito verificador inexistente"

func RunData(runIn string) (data []string, onlyRun int, err string) {

	if len(runIn) < 3 {
		return nil, 0, errRut01
	}

	if String().Contains(runIn, "-") == 0 {
		return nil, 0, errGuionRut
	}

	data = String().Split(string(runIn), "-")
	// fmt.Println("TAMAÑO", len(data), "RUT DATA -:", data)
	var e error
	onlyRun, e = strconv.Atoi(data[0])
	if e != nil {
		err = "caracteres no permitidos"
	}

	return
}
