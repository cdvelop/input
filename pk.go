package input

import "github.com/cdvelop/model"

// Primary key
// parámetro opcional:
// "show": el campo se mostrara el usuario por defecto estará oculto
func Pk(options ...string) model.Input {

	in := pk{
		Number:     Number(),
		show:       false,
		attributes: attributes{},
	}

	for _, opt := range options {
		if opt == "show" {
			in.show = true
		}
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
		HtmlTag:  in,
		Validate: in.Number,
		TestData: in.Number,
	}
}

type pk struct {
	Number model.Input
	show   bool
	attributes
}

func (p pk) Name() string {
	return "pk"
}

func (p pk) HtmlName() string {
	if p.show {
		return "number"
	}
	return "hidden"
}

// representación
func (p pk) HtmlTag(id, field_name string, allow_skip_completed bool) string {
	return p.buildHtmlTag(p.HtmlName(), p.Name(), id, field_name, allow_skip_completed)
}
