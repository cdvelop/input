package input

import (
	"strconv"
)

// options:
// title="permitido letras números - , :"
// cols="2" default 1
// rows="8" default 3
func TextArea(options ...string) *textArea {
	permitted := []rune{'%', '$', '+', '#', '-', '.', ',', ':', '(', ')'}
	var min = 5
	var max = 1000

	var info = `letras números`

	for _, p := range permitted {
		info += ` ` + string(p)
	}

	info += ` permitidos min ` + strconv.Itoa(min) + ` max ` + strconv.Itoa(max) + ` caracteres`

	new := &textArea{
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
			Letters:     true,
			Tilde:       true,
			Numbers:     true,
			BreakLine:   true,
			WhiteSpaces: true,
			Tabulation:  true,
			Characters:  permitted,
			Minimum:     min,
			Maximum:     max,
		},
	}
	new.Set(options...)

	return new
}

type textArea struct {
	attributes
	Permitted
}

func (t textArea) InputName(customName, htmlName *string) {
	if customName != nil {
		*customName = "TextArea"
	}
	if htmlName != nil {
		*htmlName = "textarea"
	}
}

func (t textArea) BuildInputHtml(id, fieldName string) string {
	return t.BuildHtmlTag("textarea", "TextArea", id, fieldName)
}
func (t textArea) ResetParameters() any {

	return &struct {
		ResetJsFuncName    string
		Enable             bool
		NotSendQueryObject bool
		Params             map[string]any
	}{
		ResetJsFuncName: "ResetTextArea",
		Enable:          true,
	}
}
