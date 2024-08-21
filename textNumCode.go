package input

import "errors"

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

func (t textNumCode) InputName(customName, htmlName *string) {
	if customName != nil {
		*customName = "TextNumCode"
	}
	if htmlName != nil {
		*htmlName = "tel"
	}
}

func (t textNumCode) BuildInputHtml(id, fieldName string) string {
	return t.BuildHtmlTag("tel", "TextNumCode", id, fieldName)
}

func (t textNumCode) ValidateInput(value string) error {

	if len(value) >= 1 {
		var ok bool
		char := value[0]

		if valid_letters[rune(char)] {
			ok = true
		}

		if valid_number[rune(char)] {
			ok = true
		}

		if !ok {
			return errors.New("no se puede comenzar con " + string(char))
		}
	}

	return t.per.Validate(value)
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
