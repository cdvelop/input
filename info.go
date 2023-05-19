package input

import "github.com/cdvelop/model"

func Info() model.Input {
	in := info{}

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

func (i info) HtmlTAG(id, field_name string, allow_skip_completed bool) string {
	return i.Value
}
