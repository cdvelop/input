package input

import "github.com/cdvelop/model"

//options:
// pattern="`^[a-zA-Z 0-9\:\.\,\+\-]{0,30}$`"
// title="permitido letras números - , :"
// cols="2" default 1
// rows="8" default 3
func TextArea(options ...string) model.Input {
	in := textArea{
		attributes: attributes{
			Rows:  `rows="3"`,
			Cols:  `cols="1"`,
			Title: `title="letras números - , : . () $ % permitidos min 2 max 1000 caracteres"`,
			// Pattern: `^[A-Za-zÑñáéíóú 0-9:$%.,+-/\\()|\n/g]{2,1000}$`,
			Oninput: `oninput="TextAreaAutoGrow(this)"`,
			// Onchange: `onchange="` + DefaultValidateFunction + `"`,
		},
		Permitted: Permitted{
			Letters:    true,
			Tilde:      true,
			Numbers:    true,
			Characters: []rune{' ', '%', '$', '+', '#', '-', '.', ',', ':', '(', ')', '\n'},
			Minimum:    2,
			Maximum:    1000,
		},
	}
	in.Set(options...)

	return model.Input{
		InputName: in.Name(),
		Tag:       &in,
		Validate:  &in,
		TestData:  &in,
	}
}

type textArea struct {
	attributes
	Permitted
}

func (t textArea) Name() string {
	return t.HtmlName()
}

func (t textArea) HtmlName() string {
	return "textarea"
}

func (t textArea) HtmlTag(id, field_name string, allow_skip_completed bool) string {
	return t.BuildHtmlTag(t.HtmlName(), t.Name(), id, field_name, allow_skip_completed)
}
