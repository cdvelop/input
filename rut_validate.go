package input

import (
	"errors"
	"strconv"
)

// validación con datos de entrada
func (r rut) ValidateInput(value string) error {

	const hidden_err = "campo invalido"

	// for _, doc := range options {
	// 	if doc == "ex" {
	// 		err := r.dni.Validate(value)
	// 		if err != nil && r.hide_typing {
	// 			return errors.New(hidden_err)
	// 		}
	// 		return err
	// 	} else {
	// 		err := r.runValidate(value)
	// 		if err != nil && r.hide_typing {
	// 			return errors.New(hidden_err)
	// 		}
	// 		return err
	// 	}
	// }

	if r.dni_mode {
		if String().Contains(value, `-`) == 0 {
			err := r.dni.Validate(value)
			if err != nil && r.hide_typing {
				return errors.New(hidden_err)
			}
			return err
		}
	}

	err := r.runValidate(value)
	if err != nil && r.hide_typing {
		return errors.New(hidden_err)
	}

	return err
}

const errCeroRut = "primer dígito no puede ser 0"

// RUT validate formato "7863697-1"
func (r rut) runValidate(rin string) error {
	data, onlyRun, err := RunData(rin)
	if err != "" {
		return errors.New(err)
	}

	if data[0][0:1] == "0" {
		return errors.New(errCeroRut)
	}

	dv := DvRut(onlyRun)

	if dv != String().ToLowerCase(data[1]) {
		return errors.New("dígito verificador " + data[1] + " inválido")
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
