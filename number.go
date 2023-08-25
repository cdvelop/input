package input

import (
	"fmt"
	"regexp"

	"github.com/cdvelop/model"
)

// options ej: data-type, data-after=" Años", pattern="^[0-9]{1,20}$",... this default
// hidden, el campo se mantendrá oculto
// title="xxx"
func Number(options ...string) model.Input {

	in := number{
		attributes: attributes{
			Title:   `title="solo valores numéricos positivos >= 0 máximo 20 char 18446744073709551615"`,
			Pattern: `^[0-9]{1,20}$`,
		},
	}
	in.Set(options...)

	out := model.Input{
		InputName: in.Name(),
		Tag:       &in,
		Validate:  &in,
		TestData:  &in,
	}

	return out
}

func Phone() model.Input {
	return Number(`pattern="^[0-9]{7,11}$"`)
}

type number struct {
	attributes
}

func (n number) Name() string {
	return "number"
}

func (n number) HtmlName() string {
	return "number"
}

func (n number) HtmlTag(id, field_name string, allow_skip_completed bool) string {
	return n.BuildHtmlTag(n.HtmlName(), n.Name(), id, field_name, allow_skip_completed)

}

// validación con datos de entrada
func (n number) ValidateField(data_in string, skip_validation bool, options ...string) bool {
	if !skip_validation {

		pvalid := regexp.MustCompile(n.Pattern)

		return pvalid.MatchString(data_in)

	} else {
		return true
	}
}

func (n number) FieldAddEventListener(field_name string) string {
	return fmt.Sprintf(`input_%v.addEventListener("input", InputValidationWithPattern);`, field_name)
}

func (n number) FieldRemoveEventListener(field_name string) string {
	return fmt.Sprintf(`input_%v.removeEventListener("input", InputValidationWithPattern);`, field_name)
}

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

	if n.Pattern != "" {
		for _, num := range temp {
			if n.ValidateField(num, false) {
				out = append(out, num)
			}
		}
	} else {
		return temp
	}

	return
}

func (n number) WrongTestData() (out []string) {
	out = []string{"1-1", "-100", "h", "h1", "-1", " ", "", "#", "& ", "% &"}

	return
}
