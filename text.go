package input

import (
	"regexp"
	"strings"

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
		Object: model.Object{
			ApiHandler: model.ApiHandler{
				Name: in.Name(),
			},
			Css:         nil,
			JsGlobal:    nil,
			JsFunctions: nil,
			JsListeners: nil,
		},
		Tag:      in,
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

func (t text) HtmlTag(id, field_name string, allow_skip_completed bool) string {
	return t.attributes.BuildHtmlTag(t.HtmlName(), t.Name(), id, field_name, allow_skip_completed)
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
func (t text) GoodTestData() (out []string) {

	first_name := []string{"Maria", "Juan", "Marcela", "Luz", "Carmen", "Jose", "Octavio"}

	last_name := []string{"Soto", "Del Rosario", "Del Carmen", "Ñuñez", "Perez", "Cadiz", "Caro"}

	phrase := []string{"Dr. Maria Jose Diaz Cadiz", "son 4 (4 bidones)", "pc dental (1)", "equipo (4)"}

	placeholder := strings.ToLower(t.PlaceHolder)

	switch {
	case strings.Contains(placeholder, "nombre y apellido"):
		return combineStringArray(true, first_name, last_name, last_name)
	case strings.Contains(placeholder, "nombre"):
		return first_name

	case strings.Contains(placeholder, "apellido"):
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
