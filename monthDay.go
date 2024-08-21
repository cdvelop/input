package input

// options: "hidden": campo oculto para el usuario
func MonthDay(options ...string) *monthDay {
	new := &monthDay{
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
	new.Set(options...)

	return new
}

// formato fecha: DD-MM
type monthDay struct {
	attributes
	Permitted
}

func (monthDay) InputName(customName, htmlName *string) {
	if customName != nil {
		*customName = "MonthDay"
	}
	if htmlName != nil {
		*htmlName = "text"
	}
}

func (m monthDay) BuildInputHtml(id, fieldName string) string {
	return m.BuildHtmlTag("text", "MonthDay", id, fieldName)
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
