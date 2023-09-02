package input

import (
	"strings"

	"github.com/cdvelop/model"
)

// SourceData() map[string]string
// options ej: "internal" = only internal contend
func CheckBox(data model.SourceData, options ...string) model.Input {
	in := check{
		Data: data,
	}

	for _, opt := range options {
		if opt == "internal" {
			in.only_internal_contend = true
		}
	}

	return model.Input{
		InputName: in.Name(),
		Tag:       &in,
		Validate:  &in,
		TestData:  &in,
	}
}

// check Box id y valor
type check struct {
	Data                  model.SourceData
	only_internal_contend bool
}

func (c check) Name() string {
	return c.HtmlName()
}

func (check) HtmlName() string {
	return "checkbox"
}

// validación con datos de entrada
func (c check) ValidateField(data_in string, skip_validation bool, options ...string) error {
	if !skip_validation {
		dataInArray := strings.Split(data_in, ",")
		for _, idkeyIn := range dataInArray {
			if _, exists := (c.Data.SourceData())[idkeyIn]; !exists {

				if idkeyIn != "" {
					return model.Error("valor", idkeyIn, "no corresponde al checkbox")

				} else {
					return model.Error("selección requerida")
				}
			}
		}

	}
	return nil
}

func (c check) GoodTestData() (out []string) {
	for k := range c.Data.SourceData() {
		out = append(out, k)
	}
	return
}

func (c check) WrongTestData() (out []string) {
	for _, wd := range wrong_data {
		if _, exist := c.Data.SourceData()[wd]; !exist {
			out = append(out, wd)
		}
	}
	return
}
