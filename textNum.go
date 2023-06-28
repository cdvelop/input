package input

import (
	"regexp"

	"github.com/cdvelop/model"
)

func TextNum() model.Input {
	in := textNum{
		attributes: attributes{
			Pattern: `^[A-Za-z0-9_]{5,20}$`,
			Title:   `title="texto, numero y guion bajo 5 a 20 caracteres"`,
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

// texto, numero y guion bajo 5 a 15 caracteres
type textNum struct {
	attributes
}

func (t textNum) Name() string {
	return "textnum"
}

func (t textNum) HtmlName() string {
	return "text"
}

func (t textNum) HtmlTag(id, field_name string, allow_skip_completed bool) string {
	return t.BuildHtmlTag(t.HtmlName(), t.Name(), id, field_name, allow_skip_completed)
}

// validación con datos de entrada
func (r textNum) ValidateField(data_in string, skip_validation bool) bool {
	if !skip_validation {

		pvalid := regexp.MustCompile(r.Pattern)

		return pvalid.MatchString(data_in)

	} else {
		return true
	}
}

func (t textNum) GoodTestData() (out []string) {
	out = []string{
		"pc_caja",
		"pc_20",
		"info_40",
		"pc_50",
		"son_24_botellas",
		"los_cuatro",
		"son_2_cuadros",
	}
	return
}

func (t textNum) WrongTestData() (out []string) {

	out = []string{
		"los cuatro",
		"tres",
		"et1_",
	}
	out = append(out, wrong_data...)

	return
}
