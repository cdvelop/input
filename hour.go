package input

import (
	"regexp"

	"github.com/cdvelop/model"
)

// formato 08:00
// options: min="08:00", max="17:00"
func Hour(options ...string) model.Input {
	in := hour{
		attributes: attributes{
			Title:   `title="formato hora: HH:MM"`,
			Pattern: `^([0-9]|0[0-9]|1[0-9]|2[0-3]):[0-5][0-9]$`,
		},
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

type hour struct {
	attributes
}

func (h hour) Name() string {
	return "hour"
}

func (h hour) HtmlName() string {
	return "time"
}

func (h hour) HtmlTAG(id, field_name string, allow_skip_completed bool) string {
	return h.Build(h.HtmlName(), h.Name(), id, field_name, allow_skip_completed)
}

// validación con datos de entrada
func (h hour) ValidateField(data_in string, skip_validation bool) bool { //en realidad es YYYY-MM-DD
	if !skip_validation {
		if len(data_in) > 10 {
			return false
		}

		pvalid := regexp.MustCompile(h.Pattern)

		return pvalid.MatchString(data_in)

	} else {
		return true
	}
}

func (h hour) GoodTestData(table_name, field_name string, random bool) (out []string) {
	out = []string{
		"23:59",
		"00:00",
		"12:00",
		"13:17",
		"21:53",
		"00:40",
		"08:30",
		"12:00",
		"15:01",
	}

	return
}

func (h hour) WrongTestData() (out []string) {
	out = []string{
		"24:00",
		"13-34",
	}
	out = append(out, wrong_data...)
	return
}
