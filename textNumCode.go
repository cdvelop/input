package input

func TextNumCode() *textNumCode {
	new := &textNumCode{
		attributes: attributes{
			// Pattern: `^[A-Za-z0-9-_]{2,15}$`,
			Title: `title="ej: V235X, 2e-45 525_45w (texto,-_, numero 2 a 15 caracteres)"`,
		},
		per: Permitted{
			Letters:    true,
			Numbers:    true,
			Characters: []rune{'_', '-'},
			Minimum:    2,
			Maximum:    15,
		},
	}

	return new
}

// texto y numero para código ej: V234
type textNumCode struct {
	attributes
	per Permitted
}

func (textNumCode) InputName() string {
	return "TextNumCode"
}

func (t textNumCode) HtmlName() string {
	return "tel"
}

func (t textNumCode) BuildContainerView(id, field_name string, allow_skip_completed bool) string {
	return t.BuildHtmlTag(t.HtmlName(), "TextNumCode", id, field_name, allow_skip_completed)
}

func (t textNumCode) ValidateField(data_in string, skip_validation bool, options ...string) (err string) {
	if !skip_validation {

		if len(data_in) >= 1 {
			var ok bool
			char := data_in[0]

			if valid_letters[rune(char)] {
				ok = true
			}

			if valid_number[rune(char)] {
				ok = true
			}

			if !ok {
				return "no se puede comenzar con " + string(char)
			}
		}

		return t.per.Validate(data_in)
	}
	return ""
}

func (t textNumCode) GoodTestData() (out []string) {

	out = []string{
		"et1",
		"12f",
		"GH03",
		"JJ10",
		"Wednesday",
		"los567",
		"677GH",
		"son_24_botellas",
	}

	return
}

func (t textNumCode) WrongTestData() (out []string) {

	out = []string{
		"los cuatro",
		"son 2 cuadros",
	}
	out = append(out, wrong_data...)

	return
}
