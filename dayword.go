package input

// formato dia DD como palabra ej. Lunes 24 Diciembre
// options: title="xxx"
func DayWord(options ...string) *dayWord {
	new := &dayWord{
		month: MonthDay(),
		attributes: attributes{
			DataSet: `data-spanish=""`,
			// Pattern: `^[0-9]{2,2}$`,
		},
	}
	new.Set(options...)

	return new
}

type dayWord struct {
	month *monthDay
	attributes
}

func (dayWord) InputName() string {
	return "DayWord"
}

func (d dayWord) HtmlName() string {
	return "text"
}

func (d dayWord) ValidateField(data_in string, skip_validation bool, options ...string) error {
	return d.month.ValidateField(data_in, skip_validation, options...)
}

func (d dayWord) GoodTestData() (out []string) {
	return d.month.GoodTestData()
}

func (d dayWord) WrongTestData() (out []string) {
	return d.month.WrongTestData()
}

func (d dayWord) BuildContainerView(id, field_name string, allow_skip_completed bool) string {
	tag := `<label class="date-spanish">`
	tag += d.BuildHtmlTag(d.HtmlName(), "DayWord", id, field_name, allow_skip_completed)
	tag += `</label>`
	return tag
}
