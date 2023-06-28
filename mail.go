package input

import (
	"regexp"

	"github.com/cdvelop/model"
)

func Mail() model.Input {
	in := mail{
		attributes: attributes{
			PlaceHolder: `placeHolder="ej: mi.correo@mail.com"`,
			Pattern:     `[a-zA-Z0-9!#$%&'*_+-]([\.]?[a-zA-Z0-9!#$%&'*_+-])+@[a-zA-Z0-9]([^@&%$\/()=?¿!.,:;]|\d)+[a-zA-Z0-9][\.][a-zA-Z]{2,4}([\.][a-zA-Z]{2})?`,
		},
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

type mail struct {
	attributes
}

func (m mail) HtmlTag(id, field_name string, allow_skip_completed bool) string {
	return m.BuildHtmlTag(m.HtmlName(), m.Name(), id, field_name, allow_skip_completed)
}

func (mail) Name() string {
	return "mail"
}

func (mail) HtmlName() string {
	return "mail"
}

// validación con datos de entrada
func (m mail) ValidateField(data_in string, skip_validation bool) bool {
	if !skip_validation {

		switch data_in {
		case "email@example.com":
			return false
		}

		pvalid := regexp.MustCompile(m.Pattern)

		return pvalid.MatchString(data_in)

	} else {
		return true
	}
}

func (mail) GoodTestData() (out []string) {

	out = []string{
		"mi.correo@mail.com",
		"alguien@algunlugar.es",
		"ramon.bonachea@email.com",
		"r.bonachea@email.com",
		"laura@hellos.email.tk",
	}

	return
}

func (mail) WrongTestData() (out []string) {

	out = []string{
		"email@example.com",
		"correomao@n//.oo",
		"son_24_bcoreos",
		"email@example",
	}
	out = append(out, wrong_data...)

	return
}
