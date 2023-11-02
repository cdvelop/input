package input

import (
	"github.com/cdvelop/model"
)

func TextNum() *model.Input {
	in := textNum{
		attributes: attributes{
			// Pattern: `^[A-Za-z0-9_]{5,20}$`,
			Title: `title="texto, numero y guion bajo 5 a 20 caracteres"`,
		},
		Permitted: Permitted{
			Letters:    true,
			Numbers:    true,
			Characters: []rune{'_'},
			Minimum:    5,
			Maximum:    20,
		},
	}

	return &model.Input{
		InputName: "TextNum",
		Tag:       &in,
		Validate:  &in,
		TestData:  &in,
	}
}

// texto, numero y guion bajo 5 a 15 caracteres
type textNum struct {
	attributes
	Permitted
}

func (t textNum) HtmlName() string {
	return "text"
}

func (t textNum) BuildContainerView(id, field_name string, allow_skip_completed bool) string {
	return t.BuildHtmlTag(t.HtmlName(), "TextNum", id, field_name, allow_skip_completed)
}

func (t textNum) GoodTestData() (out []string) {
	out = []string{
		"pc_caja",
		"pc_20",
		"info_40",
		"pc_50",
		"son_24_botellas",
		"los_cuatro",
		"son_2_cuadros",
	}
	return
}

func (t textNum) WrongTestData() (out []string) {

	out = []string{
		"los cuatro",
		"tres",
		"et1_",
	}
	out = append(out, wrong_data...)

	return
}
