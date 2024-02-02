package input

// parámetros opcionales:
// "hidden" si se vera oculto o no.
func TextOnly(options ...string) *textOnly {
	new := &textOnly{
		attributes: attributes{
			PlaceHolder: `PlaceHolder="Solo texto permitido min 3 max 50 caracteres"`,
		},
		Permitted: Permitted{
			Letters:    true,
			Minimum:    3,
			Maximum:    50,
			Characters: []rune{' '},
		},
	}
	for _, opt := range options {
		if opt == "hidden" {
			new.hidden = true
		}
	}

	return new
}

type textOnly struct {
	attributes
	hidden bool
	Permitted
}

func (textOnly) InputName() string {
	return "TextOnly"
}

func (t textOnly) HtmlName() string {
	if t.hidden {
		return "hidden"
	}
	return "text"
}

func (t textOnly) BuildContainerView(id, field_name string, allow_skip_completed bool) string {
	return t.BuildHtmlTag(t.HtmlName(), "TextOnly", id, field_name, allow_skip_completed)
}

func (t textOnly) GoodTestData() (out []string) {

	out = []string{
		"Ñuñez perez",
		"juli",
		"luz",
		"hola que tal",
		"Wednesday",
		"lost",
	}

	return
}

func (t textOnly) WrongTestData() (out []string) {

	out = []string{
		"Dr. xxx 788",
		"peréz. del jkjk",
		"los,true, vengadores",
	}
	out = append(out, wrong_data...)

	return
}
