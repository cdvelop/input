package input

import (
	"fmt"
	"regexp"

	"github.com/cdvelop/model"
)

func Date() model.Input {
	in := date{
		pattern: `[0-9]{4}-(0[1-9]|1[012])-(0[1-9]|1[0-9]|2[0-9]|3[01])`,
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

// formato fecha: DD-MM-YYYY
type date struct {
	pattern string
}

func (d date) Name() string {
	return "date"
}

func (d date) HtmlName() string {
	return "date"
}

const titleDateInfo = "formato fecha: DD-MM-YYYY"

func (d date) HtmlTAG(id, field_name string, allow_skip_completed bool) string {

	var valide string
	if !allow_skip_completed { //si cadena de validación no esta vacía
		valide = ` pattern="` + d.pattern + `" required`
	}

	tag := fmt.Sprintf(`<input`+id+`type="%v" name="%v" title="%v" %v>`, d.HtmlName(),
		field_name, titleDateInfo, valide)
	return tag
}

// validación con datos de entrada
func (d date) ValidateField(data_in string, skip_validation bool) bool {
	if !skip_validation {
		if len(data_in) > 10 {
			return false
		}

		pvalid := regexp.MustCompile(d.pattern)

		return pvalid.MatchString(data_in)

	} else {
		return true
	}
}

func (d date) GoodTestData(table_name, field_name string, random bool) (out []string) {
	out = []string{
		"2002-01-03",
		"1998-02-01",
		"1999-03-08",
		"2022-04-21",
		"1999-05-30",
		"2020-09-29",
		"1991-10-02",
		"2000-11-12",
		"1993-12-15",
	}
	return
}

func (d date) WrongTestData() (out []string) {

	out = []string{
		"21/12/1998",
		"0000-00-00",
		"31-01",
	}
	out = append(out, wrong_data...)

	return
}
