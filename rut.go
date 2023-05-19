package input

import (
	"github.com/cdvelop/model"
)

type rut struct {
	hide_typing bool
	dni_mode    bool
	attributes
}

// const patternRut = `^[0-9]+-[0-9kK]{1}$`

// parámetro opcionales:
// search-mode: modo búsqueda
// hide-typing: ocultar información  al escribir
// dni-mode: acepta documentos extranjeros
func Rut(options ...string) model.Input {
	in := rut{
		attributes: attributes{},
	}

	for _, opt := range options {
		switch opt {
		case "hide-typing":
			in.hide_typing = true
		case "dni-mode":
			in.dni_mode = true
		}
	}

	in.Autocomplete = `autocomplete="off"`

	if in.dni_mode {
		in.Title = "Documento Chileno (ch) o Extranjero (ex)"
		in.PlaceHolder = "ej: chileno (ch) 11222333-k / extranjero sin dígito ni guion (ex) 1b2334"
		in.Pattern = `^[A-Za-z0-9]{5,15}$`
	} else {
		in.Title = `title="rut sin puntos y con guion ejem.: 11222333-4"`
		in.PlaceHolder = `placeholder="ej: 11222333-4"`
		in.Oninput = `oninput="ValidateChileanDoc(this);"`
	}

	return model.Input{
		Component: model.Component{
			Name:        in.Name(),
			CssGlobal:   in,
			CssPrivate:  nil,
			JsGlobal:    in,
			JsPrivate:   in,
			JsListeners: in,
		},
		Build:    in,
		Validate: in,
		TestData: in,
	}
}

func (r rut) Name() string {
	return "rut"
}

func (r rut) HtmlName() string {
	if r.hide_typing {
		return "password"
	}
	return "text"
}
