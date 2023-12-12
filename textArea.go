package input

import (
	"strconv"

	"github.com/cdvelop/model"
)

// options:
// pattern="`^[a-zA-Z 0-9\:\.\,\+\-]{0,30}$`"
// title="permitido letras números - , :"
// cols="2" default 1
// rows="8" default 3
func TextArea(options ...string) *model.Input {
	permitted := []rune{' ', '%', '$', '+', '#', '-', '.', ',', ':', '(', ')', '\n'}
	var min = 5
	var max = 1000

	var info = `letras números`

	for _, p := range permitted {
		info += ` ` + string(p)
	}

	info += ` permitidos min ` + strconv.Itoa(min) + ` max ` + strconv.Itoa(max) + ` caracteres`

	in := textArea{
		attributes: attributes{
			Rows:  `rows="3"`,
			Cols:  `cols="1"`,
			Title: `title="` + info + `"`,
			// PlaceHolder: `placeHolder="` + info + `"`,
			// Pattern: `^[A-Za-zÑñáéíóú 0-9:$%.,+-/\\()|\n/g]{2,1000}$`,
			Oninput: `oninput="TexAreaOninput(this)"`,
			// Onchange: `onchange="` + DefaultValidateFunction + `"`,
		},
		Permitted: Permitted{
			Letters:    true,
			Tilde:      true,
			Numbers:    true,
			Characters: permitted,
			Minimum:    min,
			Maximum:    max,
		},
	}
	in.Set(options...)

	return &model.Input{
		InputName: "TextArea",
		Minimum:   min,
		Maximum:   max,
		Tag:       &in,
		Validate:  &in,
		ResetParameters: &model.ResetParameters{
			CallJsFunWithParameters: model.CallJsFunWithParameters{
				FuncNameCall: "ResetTextArea",
				Enable:       true,
				AddParams:    map[string]any{},
			},
		},
		TestData: &in,
	}
}

type textArea struct {
	attributes
	Permitted
}

func (t textArea) ResetInputView() (err string) {

	return
}

func (t textArea) HtmlName() string {
	return "textarea"
}

func (t textArea) BuildContainerView(id, field_name string, allow_skip_completed bool) string {
	return t.BuildHtmlTag(t.HtmlName(), "TextArea", id, field_name, allow_skip_completed)
}
