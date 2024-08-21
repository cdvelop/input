package input

import "errors"

func Mail() *mail {
	new := &mail{
		attributes: attributes{
			PlaceHolder: `placeHolder="ej: mi.correo@mail.com"`,
			// Pattern:     `[a-zA-Z0-9!#$%&'*_+-]([\.]?[a-zA-Z0-9!#$%&'*_+-])+@[a-zA-Z0-9]([^@&%$\/()=?¿!.,:;]|\d)+[a-zA-Z0-9][\.][a-zA-Z]{2,4}([\.][a-zA-Z]{2})?`,
		},
		per: Permitted{
			Letters:    true,
			Numbers:    false,
			Characters: []rune{'@', '.', '_'},
			Minimum:    0,
			Maximum:    0,
		},
	}

	return new
}

type mail struct {
	attributes
	per Permitted
}

func (m mail) BuildInputHtml(id, fieldName string) string {
	return m.BuildHtmlTag("mail", "Mail", id, fieldName)
}

func (mail) InputName(customName, htmlName *string) {
	if customName != nil {
		*customName = "Mail"
	}
	if htmlName != nil {
		*htmlName = "mail"
	}
}

// validación con datos de entrada
func (m mail) ValidateInput(value string) error {

	if String().Contains(value, "example") != 0 {
		return errors.New(value + " es un correo de ejemplo")
	}

	parts := String().Split(value, "@")
	if len(parts) != 2 {
		return errors.New("error en @ del correo " + value)
	}

	return m.per.Validate(value)

}

func (mail) GoodTestData() (out []string) {

	out = []string{
		"mi.correo@mail.com",
		"alguien@algunlugar.es",
		"ramon.bonachea@email.com",
		"r.bonachea@email.com",
		"laura@hellos.email.tk",
	}

	return
}

func (mail) WrongTestData() (out []string) {

	out = []string{
		"email@example.com",
		"correomao@n//.oo",
		"son_24_bcoreos",
		"email@example",
	}
	out = append(out, wrong_data...)

	return
}
