package input

// parámetros opcionales:
// "hidden" si se vera oculto o no.
// placeholder="Escriba Nombre y dos apellidos"
// title="xxx"
func Text(options ...string) *text {
	new := &text{
		attributes: attributes{
			Title: `title="texto, punto,coma, paréntesis o números permitidos max. 100 caracteres"`,
			// Pattern: `^[a-zA-ZÑñ]{2,100}[a-zA-ZÑñ0-9()., ]*$`,
		},
		Permitted: Permitted{
			Letters:    true,
			Tilde:      false,
			Numbers:    true,
			Characters: []rune{' ', '.', ',', '(', ')'},
			Minimum:    2,
			Maximum:    100,
		},
	}
	new.Set(options...)

	for _, opt := range options {
		if opt == "hidden" {
			new.hidden = true
		}
	}

	return new
}

// texto,punto,coma, paréntesis o números permitidos
type text struct {
	hidden bool
	attributes
	Permitted
}

func (text) InputName() string {
	return "Text"
}

func (t text) HtmlName() string {
	if t.hidden {
		return "hidden"
	}
	return "text"
}

func (t text) BuildContainerView(id, field_name string, allow_skip_completed bool) string {
	return t.attributes.BuildHtmlTag(t.HtmlName(), "Text", id, field_name, allow_skip_completed)
}

// options: first_name,last_name, phrase
func (t text) GoodTestData() (out []string) {

	first_name := []string{"Maria", "Juan", "Marcela", "Luz", "Carmen", "Jose", "Octavio"}

	last_name := []string{"Soto", "Del Rosario", "Del Carmen", "Ñuñez", "Perez", "Cadiz", "Caro"}

	phrase := []string{"Dr. Maria Jose Diaz Cadiz", "son 4 (4 bidones)", "pc dental (1)", "equipo (4)"}

	placeholder := String().ToLowerCase(t.PlaceHolder)

	switch {
	case String().Contains(placeholder, "nombre y apellido") != 0:

		return permutation(first_name, last_name)
	case String().Contains(placeholder, "nombre") != 0:
		return first_name

	case String().Contains(placeholder, "apellido") != 0:
		return last_name

	default:
		out = append(out, phrase...)
		out = append(out, first_name...)
		out = append(out, last_name...)
	}

	return
}

func (text) WrongTestData() (out []string) {

	out = []string{
		"peréz del rozal",
		" estos son \\n los podria",
	}

	out = append(out, wrong_data...)

	return
}
