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
			Title:   `title="Campo Informativo"`,
			Pattern: `[0-9]{4}-(0[1-9]|1[012])-(0[1-9]|1[0-9]|2[0-9]|3[01])`,
			// Onkeyup:  `onkeyup="DateAgeChange(this)"`,
			Onchange: `oninput="DateAgeChange(this)"`,
		},
	}
	in.Set(options...)

	return model.Input{
		InputName: in.Name(),
		Tag:       &in,
		Validate:  &in,
		TestData:  &in,
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

func (d dateAge) HtmlTag(id, field_name string, allow_skip_completed bool) string {

	tag := `<label class="age-number"><input data-name="age-number" type="number" min="0" max="150" oninput="AgeInputChange(this)" title="Campo Informativo"></label>`

	tag += `<label class="age-date">`

	tag += d.BuildHtmlTag(d.HtmlName(), d.Name(), id, field_name, allow_skip_completed)

	tag += `</label>`

	return tag
}

func (d dateAge) ValidateField(data_in string, skip_validation bool, options ...string) bool { //en realidad es YYYY-MM-DD
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

func (d dateAge) GoodTestData() (out []string) {
	return Date().TestData.GoodTestData()
}

func (d dateAge) WrongTestData() (out []string) {
	return Date().TestData.WrongTestData()
}
