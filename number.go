package input

import (
	"strconv"

	"github.com/cdvelop/model"
)

// options ej: data-type, data-after=" Años"
// hidden, el campo se mantendrá oculto
// title="xxx"
// for phone ej: `min="7"`, `max="11"`
func Number(options ...string) *model.Input {

	in := number{
		attributes: attributes{
			Title: `title="solo valores numéricos positivos >= 0 máximo 20 char 18446744073709551615"`,
			// Pattern: `^[0-9]{1,20}$`,
		},
		Permitted: Permitted{
			Numbers: true,
			Minimum: 1,
			Maximum: 20,
		},
	}
	in.Set(options...)

	if in.Min != "" {
		in.Minimum, _ = strconv.Atoi(in.Min)
	}

	if in.Max != "" {
		in.Maximum, _ = strconv.Atoi(in.Max)
	}

	return &model.Input{
		InputName: "Number",
		Tag:       &in,
		Validate:  &in,
		TestData:  &in,
	}
}

func Phone() *model.Input {
	// Phone `pattern="^[0-9]{7,11}$"`
	return Number(`min="7"`, `max="11"`)
}

type number struct {
	attributes
	Permitted
}

func (n number) HtmlName() string {
	return "number"
}

func (n number) BuildContainerView(id, field_name string, allow_skip_completed bool) string {
	return n.BuildHtmlTag(n.HtmlName(), "Number", id, field_name, allow_skip_completed)

}

// func (n number) FieldAddEventListener(field_name string) string {
// 	return fmt.Sprintf(`input_%v.addEventListener("input", InputValidationWithPattern);`, field_name)
// }

// func (n number) FieldRemoveEventListener(field_name string) string {
// 	return fmt.Sprintf(`input_%v.removeEventListener("input", InputValidationWithPattern);`, field_name)
// }

func (n number) GoodTestData() (out []string) {

	temp := []string{
		"56988765432",
		"1234567",
		"0",
		"123456789",
		"100",
		"5000",
		"423456789",
		"31",
		"523756789",
		"10000232326263727",
		"29",
		"923726789",
		"3234567",
		"823456789",
		"29",
	}

	for _, v := range temp {
		if len(v) >= n.Minimum && len(v) <= n.Maximum {
			out = append(out, v)
		}
	}

	return
}

func (n number) WrongTestData() (out []string) {
	out = []string{"1-1", "-100", "h", "h1", "-1", " ", "", "#", "& ", "% &"}

	return
}
