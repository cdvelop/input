package input

import (
	"regexp"

	"github.com/cdvelop/model"
)

func TextNumCode() model.Input {
	in := textNumCode{
		attributes: attributes{
			Pattern: `^[A-Za-z0-9-_]{2,15}$`,
			Title:   `title="ej: V235X, 2e-45 525_45w (texto,-_, numero 2 a 15 caracteres)"`,
		},
	}

	return model.Input{
		InputName: in.Name(),
		Tag:       &in,
		Validate:  &in,
		TestData:  &in,
	}
}

// texto y numero para código ej: V234
type textNumCode struct {
	attributes
}

func (t textNumCode) Name() string {
	return "textnumcode"
}

func (t textNumCode) HtmlName() string {
	return "tel"
}

func (t textNumCode) HtmlTag(id, field_name string, allow_skip_completed bool) string {
	return t.BuildHtmlTag(t.HtmlName(), t.Name(), id, field_name, allow_skip_completed)
}

// validación con datos de entrada
func (t textNumCode) ValidateField(data_in string, skip_validation bool) bool {
	if !skip_validation {
		switch data_in {
		case "-1":
			return false

		}

		pvalid := regexp.MustCompile(t.Pattern)

		return pvalid.MatchString(data_in)

	} else {
		return true
	}
}

func (t textNumCode) GoodTestData() (out []string) {

	out = []string{
		"et1",
		"12f",
		"GH03",
		"JJ10",
		"Wednesday",
		"los567",
		"677GH",
		"son_24_botellas",
	}

	return
}

func (t textNumCode) WrongTestData() (out []string) {

	out = []string{
		"los cuatro",
		"son 2 cuadros",
	}
	out = append(out, wrong_data...)

	return
}
