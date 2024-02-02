package input

// formato 08:00
// options: min="08:00", max="17:00"
func Hour(options ...string) *hour {
	new := &hour{
		attributes: attributes{
			Title: `title="formato hora: HH:MM"`,
			// Pattern: `^([0-9]|0[0-9]|1[0-9]|2[0-3]):[0-5][0-9]$`,
		},
		per: Permitted{
			Numbers:    true,
			Characters: []rune{':'},
			Minimum:    5,
			Maximum:    5,
		},
	}
	new.Set(options...)

	return new
}

type hour struct {
	attributes
	per Permitted
}

func (hour) InputName() string {
	return "Hour"
}

func (hour) HtmlName() string {
	return "time"
}

func (h hour) BuildContainerView(id, field_name string, allow_skip_completed bool) string {
	return h.BuildHtmlTag(h.HtmlName(), "Hour", id, field_name, allow_skip_completed)
}

func (h hour) ValidateField(data_in string, skip_validation bool, options ...string) (err string) {
	if !skip_validation {

		if len(data_in) >= 2 && data_in[0] == '2' && data_in[1] == '4' {
			return "la hora 24 no existe"
		}

		return h.per.Validate(data_in)

	}
	return ""
}

func (h hour) GoodTestData() (out []string) {
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
