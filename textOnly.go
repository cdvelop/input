package input

import (
	"regexp"

	"github.com/cdvelop/model"
)

// parámetros opcionales:
// "hidden" si se vera oculto o no.
func TextOnly(options ...string) model.Input {
	in := textOnly{
		attributes: attributes{
			PlaceHolder: `PlaceHolder="Solo texto permitido min 3 max 50 caracteres"`,
			Pattern:     `^[A-Za-zÑñáéíóú ]{3,50}$`,
		},
	}
	for _, opt := range options {
		if opt == "hidden" {
			in.hidden = true
		}
	}

	return model.Input{
		Name:     in.Name(),
		Tag:      in,
		Validate: in,
		TestData: in,
	}
}

type textOnly struct {
	attributes
	hidden bool
}

func (t textOnly) Name() string {
	return "text"
}

func (t textOnly) HtmlName() string {
	if t.hidden {
		return "hidden"
	}
	return "text"
}

func (t textOnly) HtmlTag(id, field_name string, allow_skip_completed bool) string {
	return t.BuildHtmlTag(t.HtmlName(), t.Name(), id, field_name, allow_skip_completed)
}

// validación con datos de entrada
func (t textOnly) ValidateField(data_in string, skip_validation bool) bool {
	if !skip_validation {
		pvalid := regexp.MustCompile(t.Pattern)

		return pvalid.MatchString(data_in)

	} else {
		return true
	}
}

func (t textOnly) GoodTestData() (out []string) {

	out = []string{
		"Ñuñez perez",
		"juli",
		"luz",
		"hola que tal",
		"Wednesday",
		"lost",
	}

	return
}

func (t textOnly) WrongTestData() (out []string) {

	out = []string{
		"Dr. xxx 788",
		"peréz. del jkjk",
		"los,true, vengadores",
	}
	out = append(out, wrong_data...)

	return
}
