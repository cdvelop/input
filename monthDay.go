package input

import (
	"github.com/cdvelop/model"
)

// options: "hidden": campo oculto para el usuario
func MonthDay(options ...string) model.Input {
	in := monthDay{
		attributes: attributes{
			// Pattern: `^[0-9]{2,2}$`,
		},
		Permitted: Permitted{
			Numbers:    true,
			Characters: []rune{},
			Minimum:    2,
			Maximum:    2,
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

// formato fecha: DD-MM
type monthDay struct {
	attributes
	Permitted
}

func (d monthDay) Name() string {
	return "monthday"
}

func (d monthDay) HtmlName() string {
	return "text"
}

func (m monthDay) HtmlTag(id, field_name string, allow_skip_completed bool) string {
	return m.BuildHtmlTag(m.HtmlName(), m.Name(), id, field_name, allow_skip_completed)
}

func (m monthDay) GoodTestData() (out []string) {

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
