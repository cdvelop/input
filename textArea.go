package input

import "github.com/cdvelop/model"

// pattern ej: `^[a-zA-Z 0-9\:\.\,\+\-]{0,30}$`
// info ej: `permitido letras números - , :
// cols ej: 1,2 ... 0 = default 1
// rows ej; 4,6 ... 0 = default 3
func TextArea(pattern, info string, cols, rows uint8) model.Input {
	in := textArea{}

	if pattern != "" {
		in.pattern = pattern
	} else {
		in.pattern = `^[A-Za-zÑñáéíóú 0-9:$%.,+-/\\()|\n/g]{2,1000}$`
	}
	if info != "" {
		in.info = info
	} else {
		in.info = `letras números - , : . () $ % permitidos min 2 max 1000 caracteres`
	}

	if rows != 0 && cols != 0 {
		in.Rows = rows
		in.Cols = cols
	} else {
		in.Cols = 1
		in.Rows = 3
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

type textArea struct {
	info    string
	pattern string
	Rows    byte //filas ej 4,5,6
	Cols    byte //columnas ej 50,80
}

func (t textArea) Name() string {
	return t.HtmlName()
}

func (t textArea) HtmlName() string {
	return "textarea"
}
