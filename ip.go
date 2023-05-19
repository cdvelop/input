package input

import (
	"regexp"

	"github.com/cdvelop/model"
)

// dirección ip valida campos separados por puntos
func Ip() model.Input {
	in := ip{
		attributes: attributes{
			Title:   `title="dirección ip valida campos separados por puntos ej 192.168.0.8"`,
			Pattern: `^((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$`,
		},
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

type ip struct {
	attributes
}

func (i ip) Name() string {
	return "ip"
}

func (i ip) HtmlName() string {
	return "text"
}

func (i ip) HtmlTAG(id, field_name string, allow_skip_completed bool) string {

	return i.Build(i.HtmlName(), i.Name(), id, field_name, allow_skip_completed)

}

// validación con datos de entrada
func (i ip) ValidateField(data_in string, skip_validation bool) bool {
	if !skip_validation {

		if data_in == "0.0.0.0" {
			return false
		}

		pvalid := regexp.MustCompile(i.Pattern)

		return pvalid.MatchString(data_in)

	} else {
		return true
	}
}
