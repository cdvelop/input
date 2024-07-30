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

func (dateAge) InputName() string {
	return "DateAge"
}

func (d dateAge) HtmlName() string {
	return "date"
}

func (d dateAge) BuildContainerView(id, field_name string, allow_skip_completed bool) string {

	tag := `<label class="age-number"><input data-name="age-number" type="number" min="0" max="150" oninput="AgeInputChange(this)" title="Campo Informativo"></label>`

	tag += `<label class="age-date">`

	tag += d.BuildHtmlTag(d.HtmlName(), "DateAge", id, field_name, allow_skip_completed)

	tag += `</label>`

	return tag
}

func (d dateAge) ValidateField(data_in string, skip_validation bool, options ...string) error { //en realidad es YYYY-MM-DD
	if !skip_validation {
		return d.day.CheckDateExists(data_in)

	}
	return nil
}

func (d dateAge) GoodTestData() (out []string) {
	return d.day.GoodTestData()
}

func (d dateAge) WrongTestData() (out []string) {
	return d.day.WrongTestData()
}
