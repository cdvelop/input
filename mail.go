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

func (m mail) BuildContainerView(id, field_name string, allow_skip_completed bool) string {
	return m.BuildHtmlTag(m.HtmlName(), "Mail", id, field_name, allow_skip_completed)
}

func (mail) InputName() string {
	return "Mail"
}

func (mail) HtmlName() string {
	return "mail"
}

// validación con datos de entrada
func (m mail) ValidateField(data_in string, skip_validation bool, options ...string) error {
	if !skip_validation {

		if String().Contains(data_in, "example") != 0 {
			return errors.New(data_in + " es un correo de ejemplo")
		}

		parts := String().Split(data_in, "@")
		if len(parts) != 2 {
			return errors.New("error en @ del correo " + data_in)
		}

		return m.per.Validate(data_in)

	}

	return nil
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
