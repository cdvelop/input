package input

import (
	"github.com/cdvelop/model"
)

// formato dia DD como palabra ej. Lunes 24 Diciembre
// options: title="xxx"
func DayWord(options ...string) *model.Input {
	in := dayWord{
		attributes: attributes{
			DataSet: `data-spanish=""`,
			// Pattern: `^[0-9]{2,2}$`,
		},
	}
	in.Set(options...)

	month_day := MonthDay()

	return &model.Input{
		InputName: "DayWord",
		Tag:       in,
		Validate:  month_day,
		TestData:  month_day,
	}
}

type dayWord struct {
	attributes
}

func (d dayWord) HtmlName() string {
	return "text"
}

func (d dayWord) BuildContainerView(id, field_name string, allow_skip_completed bool) string {
	tag := `<label class="date-spanish">`
	tag += d.BuildHtmlTag(d.HtmlName(), "DayWord", id, field_name, allow_skip_completed)
	tag += `</label>`
	return tag
}
