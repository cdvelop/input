package input

import (
	"regexp"

	"github.com/cdvelop/model"
)

// formato dia DD como palabra ej. Lunes 24 Diciembre
// options: title="xxx"
func DayWord(options ...string) model.Input {
	in := dayWord{
		attributes: attributes{
			DataSet: `data-spanish=""`,
			Pattern: `^[0-9]{2,2}$`,
		},
	}
	in.Set(options...)

	return model.Input{
		InputName: in.Name(),
		Tag:       in,
		Validate:  in,
		TestData:  in,
	}
}

type dayWord struct {
	attributes
}

func (d dayWord) Name() string {
	return "dayword"
}

func (d dayWord) HtmlName() string {
	return "text"
}

func (d dayWord) HtmlTag(id, field_name string, allow_skip_completed bool) string {
	tag := `<label class="date-spanish">`
	tag += d.BuildHtmlTag(d.HtmlName(), d.Name(), id, field_name, allow_skip_completed)
	tag += `</label>`
	return tag
}

// validación con datos de entrada
func (d dayWord) ValidateField(data_in string, skip_validation bool) bool { //en realidad es YYYY-MM-DD
	if !skip_validation {

		pvalid := regexp.MustCompile(d.Pattern)

		return pvalid.MatchString(data_in)

	} else {
		return true
	}
}

func (dayWord) GoodTestData() (out []string) {
	return MonthDay().GoodTestData()
}

func (dayWord) WrongTestData() (out []string) {
	return MonthDay().WrongTestData()
}
