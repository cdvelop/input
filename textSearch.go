package input

func TextSearch() *textSearch {

	return &textSearch{
		attributes: attributes{
			// Pattern: `^[a-zA-ZÑñ0-9- ]{2,20}$`,
			Title: `title="letras números y guion - permitido. max 20 caracteres"`,
		},
		Permitted: Permitted{
			Letters:    true,
			Tilde:      false,
			Numbers:    true,
			Characters: []rune{'-', ' '},
			Minimum:    2,
			Maximum:    20,
		},
	}
}

type textSearch struct {
	attributes
	Permitted
}

func (t textSearch) InputName(customName, htmlName *string) {
	if customName != nil {
		*customName = "TextSearch"
	}
	if htmlName != nil {
		*htmlName = "search"
	}
}

func (t textSearch) BuildInputHtml(id, fieldName string) string {
	return t.BuildHtmlTag("search", "TextSearch", id, fieldName)
}

func (s textSearch) GoodTestData() (out []string) {
	out = []string{
		"Ñuñez perez",
		"Maria Jose Diaz",
		"12038-0",
		"1990-07-21",
		"190-07-21",
		"lost",
	}
	return
}

func (s textSearch) WrongTestData() (out []string) {
	out = []string{
		"Dr. xxx 788",
		"peréz del jkjk",
		"los,true, vengadores",
		"0", " ", "", "#", "& ", "% &", "SELECT * FROM", "=",
	}

	return
}
