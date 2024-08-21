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

func (t textOnly) InputName(customName, htmlName *string) {
	if customName != nil {
		*customName = "TextOnly"
	}
	if htmlName != nil {
		if t.hidden {
			*htmlName = "hidden"
		} else {
			*htmlName = "text"
		}
	}
}

func (t textOnly) BuildInputHtml(id, fieldName string) string {
	htmlName := "text"
	if t.hidden {
		htmlName = "hidden"
	}
	return t.BuildHtmlTag(htmlName, "TextOnly", id, fieldName)
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
