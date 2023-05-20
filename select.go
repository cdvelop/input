package input

import "github.com/cdvelop/model"

// SourceData() map[string]string
func SelecTag(data sourceData) model.Input {
	in := selecTag{
		Data: data,
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
		Build:    in,
		Validate: in,
		TestData: in,
	}
}

type selecTag struct {
	Data sourceData
}

func (s selecTag) Name() string {
	return "selecTag"
}

func (s selecTag) HtmlName() string {
	return "selecTag"
}
