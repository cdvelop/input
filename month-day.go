package input

import (
	"regexp"
	"strconv"

	"github.com/cdvelop/model"
)

// options: "hidden": campo oculto para el usuario
func MonthDay(options ...string) model.Input {
	in := monthDay{
		attributes: attributes{},
	}
	in.Set(options...)

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

const patternMonthDay = `^[0-9]{2,2}$`

// formato fecha: DD-MM
type monthDay struct {
	attributes
}

func (d monthDay) Name() string {
	return "monthday"
}

func (d monthDay) HtmlName() string {
	return "text"
}

func (m monthDay) HtmlTAG(id, field_name string, allow_skip_completed bool) string {
	return m.Build(m.HtmlName(), m.Name(), id, field_name, allow_skip_completed)
}

// validación con datos de entrada
func (m monthDay) ValidateField(data_in string, skip_validation bool) bool {
	if !skip_validation {

		intVar, err := strconv.Atoi(data_in)
		if err != nil {
			return false
		}

		if intVar > 31 {
			return false
		}

		pvalid := regexp.MustCompile(patternMonthDay)

		return pvalid.MatchString(data_in)

	} else {
		return true
	}
}

func (m monthDay) GoodTestData(table_name, field_name string, random bool) (out []string) {

	out = []string{
		"01",
		"30",
		"03",
		"22",
		"31",
		"29",
		"10",
		"12",
		"05",
	}

	return
}

func (m monthDay) WrongTestData() (out []string) {
	out = []string{
		"1-1",
		"21/12",
	}

	out = append(out, wrong_data...)

	return
}
