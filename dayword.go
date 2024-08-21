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

func (d dayWord) InputName(customName, htmlName *string) {
	if customName != nil {
		*customName = "DayWord"
	}
	if htmlName != nil {
		*htmlName = "text"
	}
}

func (d dayWord) BuildInputHtml(id, fieldName string) string {
	tag := `<label class="date-spanish">`
	tag += d.BuildHtmlTag("text", "DayWord", id, fieldName)
	tag += `</label>`
	return tag
}

func (d dayWord) ValidateInput(value string) error {
	return d.month.ValidateInput(value)
}

func (d dayWord) GoodTestData() (out []string) {
	return d.month.GoodTestData()
}

func (d dayWord) WrongTestData() (out []string) {
	return d.month.WrongTestData()
}
