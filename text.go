package input

import (
	"regexp"

	"github.com/cdvelop/model"
)

// parámetros opcionales:
// "hidden" si se vera oculto o no.
// placeholder="Escriba Nombre y dos apellidos"
// title="xxx"
func Text(options ...string) model.Input {
	in := text{
		attributes: attributes{
			Title:   `title="texto, punto,coma, paréntesis o números permitidos max. 100 caracteres"`,
			Pattern: `^[a-zA-ZÑñ]{2,100}[a-zA-ZÑñ0-9()., ]*$`,
		},
	}
	in.Set(options...)

	for _, opt := range options {
		if opt == "hidden" {
			in.hidden = true
		}
	}

	return model.Input{
		Component: model.Component{
			Name:        in.Name(),
			CssGlobal:   nil,
			CssPrivate:  nil,
			JsGlobal:    nil,
			JsPrivate:   nil,
			JsListeners: nil,
		},
		Build:    in,
		Validate: in,
		TestData: in,
	}
}

// texto,punto,coma, paréntesis o números permitidos
type text struct {
	hidden bool
	attributes
}

func (text) Name() string {
	return "text"
}

func (t text) HtmlName() string {
	if t.hidden {
		return "hidden"
	}
	return "text"
}

func (t text) HtmlTAG(id, field_name string, allow_skip_completed bool) string {
	return t.Build(t.HtmlName(), t.Name(), id, field_name, allow_skip_completed)
}

// validación con datos de entrada
func (t text) ValidateField(data_in string, skip_validation bool) bool {
	if !skip_validation {

		pvalid := regexp.MustCompile(t.Pattern)

		return pvalid.MatchString(data_in)

	} else {
		return true
	}
}

// options: first_name,last_name, phrase
func (text) GoodTestData(table_name, field_name string, random bool) (out []string) {

	first_name := []string{"Maria", "Juan", "Marcela", "Luz", "Carmen", "Jose", "Octavio"}

	last_name := []string{"Soto", "Del Rosario", "Del Carmen", "Ñuñez", "Perez", "Cadiz", "Caro"}

	phrase := []string{"Dr. Maria Jose Diaz Cadiz", "son 4 (4 bidones)", "pc dental (1)", "equipo (4)"}

	switch field_name {
	case table_name + "_name":
		return combineStringArray(true, first_name, last_name, last_name)
	case "first_name":
		return first_name

	case "last_name":
		return last_name

	default:
		out = append(out, phrase...)
		out = append(out, first_name...)
		out = append(out, last_name...)
	}

	return
}

func (text) WrongTestData() (out []string) {

	out = []string{
		"peréz del rozal",
		" estos son \\n los podria",
	}

	out = append(out, wrong_data...)

	return
}
