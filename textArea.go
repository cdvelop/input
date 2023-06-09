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
			Rows:    `rows="3"`,
			Cols:    `cols="1"`,
			Title:   `title="letras números - , : . () $ % permitidos min 2 max 1000 caracteres"`,
			Pattern: `^[A-Za-zÑñáéíóú 0-9:$%.,+-/\\()|\n/g]{2,1000}$`,
			Oninput: `oninput="TextAreaAutoGrow(this)"`,
			Onkeyup: `onkeyup="` + DefaultValidateFunction + `"`,
		},
	}
	in.Set(options...)

	return model.Input{
		Object: model.Object{
			ApiHandler: model.ApiHandler{
				Name: in.Name(),
			},
			Css:         nil,
			JsGlobal:    in,
			JsFunctions: nil,
			JsListeners: nil,
		},
		Tag:      in,
		Validate: in,
		TestData: in,
	}
}

type textArea struct {
	attributes
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
func (t textArea) JsGlobal() string {
	return `function TextAreaAutoGrow(input) {
		input.style.height = "5px";
		input.style.height = (input.scrollHeight) + "px";
	};`
}
