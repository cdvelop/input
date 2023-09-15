package input

import (
	"strconv"
	"strings"

	"github.com/cdvelop/model"
)

func Date() *model.Input {
	in := date{
		attributes: attributes{
			Title: `title="formato fecha: DD-MM-YYYY"`,
			// Pattern: `[0-9]{4}-(0[1-9]|1[012])-(0[1-9]|1[0-9]|2[0-9]|3[01])`,
		},
	}

	return &model.Input{
		InputName: in.Name(),
		Tag:       &in,
		Validate:  &in,
		TestData:  &in,
	}
}

// formato fecha: DD-MM-YYYY
type date struct {
	attributes
}

func (d date) Name() string {
	return "Date"
}

func (d date) HtmlName() string {
	return "date"
}

func (d date) HtmlTag(id, field_name string, allow_skip_completed bool) string {
	return d.BuildHtmlTag(d.HtmlName(), d.Name(), id, field_name, allow_skip_completed)
}

// validación con datos de entrada
func (d date) ValidateField(data_in string, skip_validation bool, options ...string) error {
	if !skip_validation {
		return validateDate(data_in)
	}
	return nil
}

func validateDate(data_in string) error {

	if data_in == "0000-00-00" {
		return model.Error("fecha ejemplo no válida")
	}
	// Dividir la cadena en partes separadas por "-"
	parts := strings.Split(data_in, "-")

	// Verificar si hay tres partes (año, mes y día)
	if len(parts) != 3 {
		return model.Error("formato fecha no válido")
	}

	// Verificar si cada parte es un número válido
	for _, part := range parts {
		// Intentar convertir la parte en un número entero
		num, err := strconv.Atoi(part)
		if err != nil {
			return model.Error("No es un número válido")
		}

		// Verificar los rangos para año, mes y día
		if part == parts[0] && (num < 1000 || num > 9999) {
			return model.Error("Año no válido")
		} else if part == parts[1] && (num < 1 || num > 12) {
			return model.Error("Mes no válido")
		} else if part == parts[2] && (num < 1 || num > 31) {
			return model.Error("Día no válido")
		}
	}

	return nil
}

func (d date) GoodTestData() (out []string) {
	out = []string{
		"2002-01-03",
		"1998-02-01",
		"1999-03-08",
		"2022-04-21",
		"1999-05-30",
		"2020-09-29",
		"1991-10-02",
		"2000-11-12",
		"1993-12-15",
	}
	return
}

func (d date) WrongTestData() (out []string) {

	out = []string{
		"21/12/1998",
		"0000-00-00",
		"31-01",
	}
	out = append(out, wrong_data...)

	return
}
