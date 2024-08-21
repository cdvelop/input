package input

// formato fecha: DD-MM-YYYY
// options: `title="xxx"`
func DateAge(options ...string) *dateAge {
	new := dateAge{
		attributes: attributes{
			Title: `title="formato fecha: DD-MM-YYYY"`,
			// Pattern: `[0-9]{4}-(0[1-9]|1[012])-(0[1-9]|1[0-9]|2[0-9]|3[01])`,
			// Onkeyup:  `onkeyup="DateAgeChange(this)"`,
			Onchange: `onchange="DateAgeChange(this)"`,
		},
		day: Date(),
	}
	new.Set(options...)

	return &new
}

type dateAge struct {
	attributes
	day *date
}

func (d dateAge) InputName(customName, htmlName *string) {
	if customName != nil {
		*customName = "DateAge"
	}
	if htmlName != nil {
		*htmlName = "date"
	}
}

func (d dateAge) BuildInputHtml(id, fieldName string) string {

	tag := `<label class="age-number"><input data-name="age-number" type="number" min="0" max="150" oninput="AgeInputChange(this)" title="Campo Informativo"></label>`

	tag += `<label class="age-date">`

	tag += d.BuildHtmlTag("date", "DateAge", id, fieldName)

	tag += `</label>`

	return tag
}

func (d dateAge) ValidateInput(value string) error { //en realidad es YYYY-MM-DD
	return d.day.CheckDateExists(value)
}

func (d dateAge) GoodTestData() (out []string) {
	return d.day.GoodTestData()
}

func (d dateAge) WrongTestData() (out []string) {
	return d.day.WrongTestData()
}
