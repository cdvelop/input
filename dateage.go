package input

import (
	"regexp"

	"github.com/cdvelop/model"
)

// formato fecha: DD-MM-YYYY
// options: `title="xxx"`
func DateAge(options ...string) model.Input {
	in := dateAge{
		attributes: attributes{
			Pattern: `[0-9]{4}-(0[1-9]|1[012])-(0[1-9]|1[0-9]|2[0-9]|3[01])`,
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

type dateAge struct {
	attributes
}

func (d dateAge) Name() string {
	return "dateage"
}

func (d dateAge) HtmlName() string {
	return "date"
}

func (d dateAge) ValidateField(data_in string, skip_validation bool) bool { //en realidad es YYYY-MM-DD
	if !skip_validation {
		if len(data_in) > 10 {
			return false
		}

		pvalid := regexp.MustCompile(d.Pattern)

		return pvalid.MatchString(data_in)

	} else {
		return true
	}
}

func (d dateAge) GoodTestData(table_name, field_name string, random bool) (out []string) {
	return Date().TestData.GoodTestData("", "", random)
}

func (d dateAge) WrongTestData() (out []string) {
	return Date().TestData.WrongTestData()
}
