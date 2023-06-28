package input

import "github.com/cdvelop/model"

func Info(value string) model.Input {
	in := info{
		Value: value,
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
		Validate: nil,
		TestData: nil,
	}
}

// input de carácter informativo
type info struct {
	Value string //valor a mostrar
}

func (i info) Name() string {
	return "info"
}

func (i info) HtmlName() string {
	return "text"
}

func (i info) HtmlTag(id, field_name string, allow_skip_completed bool) string {
	return i.Value
}
