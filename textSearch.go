package input

import (
	"github.com/cdvelop/model"
)

func TextSearch() *model.Input {
	in := textSearch{
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

	return &model.Input{
		InputName: "TextSearch",
		Tag:       &in,
		Validate:  &in,
		TestData:  &in,
	}
}

type textSearch struct {
	attributes
	Permitted
}

func (t textSearch) HtmlName() string {
	return "search"
}

func (t textSearch) BuildContainerView(id, field_name string, allow_skip_completed bool) string {
	return t.BuildHtmlTag(t.HtmlName(), "TextSearch", id, field_name, allow_skip_completed)

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
