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
		Tag:      in,
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

func (i ip) HtmlTag(id, field_name string, allow_skip_completed bool) string {
	return i.BuildHtmlTag(i.HtmlName(), i.Name(), id, field_name, allow_skip_completed)

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

func (i ip) GoodTestData() (out []string) {

	out = []string{
		"120.1.3.206",
		"195.145.149.184",
		"179.183.230.16",
		"253.70.9.26",
		"215.35.117.51",
		"212.149.243.253",
		"126.158.214.250",
		"49.122.253.195",
		"53.218.195.25",
		"190.116.115.117",
		"115.186.149.240",
		"163.95.226.221",
	}

	return
}

func (i ip) WrongTestData() (out []string) {
	out = []string{
		"0.0.0.0",
		"192.168.1.1.8",
	}
	out = append(out, wrong_data...)
	return
}
