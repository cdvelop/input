package input

import (
	"regexp"

	"github.com/cdvelop/model"
)

func TextSearch() model.Input {
	in := textSearch{
		attributes: attributes{
			Pattern: `^[a-zA-ZÑñ0-9- ]{2,20}$`,
			Title:   `title="letras números y guion - permitido. max 20 caracteres"`,
		},
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

type textSearch struct {
	attributes
}

func (t textSearch) Name() string {
	return "textsearch"
}

func (t textSearch) HtmlName() string {
	return "search"
}

func (t textSearch) HtmlTAG(id, field_name string, allow_skip_completed bool) string {

	return t.Build(t.HtmlName(), t.Name(), id, field_name, allow_skip_completed)

}

// validación con datos de entrada
func (t textSearch) ValidateField(data_in string, skip_validation bool) bool {
	if !skip_validation {

		pvalid := regexp.MustCompile(t.Pattern)

		return pvalid.MatchString(data_in)

	} else {
		return true
	}
}

func (s textSearch) GoodTestData(table_name, field_name string, random bool) (out []string) {
	out = []string{
		"Ñuñez perez",
		"Maria Jose Diaz",
		"12038-0",
		"1990-07-21",
		"190-07-21",
		"lost",
	}
	return
}

func (s textSearch) WrongTestData() (out []string) {
	out = []string{
		"Dr. xxx 788",
		"peréz del jkjk",
		"los,true, vengadores",
		"0", " ", "", "#", "& ", "% &", "SELECT * FROM", "=",
	}

	return
}
