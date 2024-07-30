package input

func Date() *date {

	return &date{
		attributes: attributes{
			Title: `title="formato fecha: DD-MM-YYYY"`,
			// Pattern: `[0-9]{4}-(0[1-9]|1[012])-(0[1-9]|1[0-9]|2[0-9]|3[01])`,
		},
	}
}

// formato fecha: DD-MM-YYYY
type date struct {
	attributes
}

func (date) InputName() string {
	return "Date"
}

func (d date) HtmlName() string {
	return "date"
}

func (d date) BuildContainerView(id, field_name string, allow_skip_completed bool) string {
	return d.BuildHtmlTag(d.HtmlName(), "Date", id, field_name, allow_skip_completed)
}

// validación con datos de entrada
func (d date) ValidateField(data_in string, skip_validation bool, options ...string) error {
	if !skip_validation {
		return d.CheckDateExists(data_in)
	}
	return nil
}

func (d date) GoodTestData() (out []string) {
	out = []string{
		"2002-01-03",
		"1998-02-01",
		"1999-03-08",
		"2022-04-21",
		"1999-05-30",
		"2020-09-29",
		"1991-10-02",
		"2000-11-12",
		"1993-12-15",
	}
	return
}

func (d date) WrongTestData() (out []string) {

	out = []string{
		"21/12/1998",
		"0000-00-00",
		"31-01",
	}
	out = append(out, wrong_data...)

	return
}
