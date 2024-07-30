package input

import "errors"

// name ej: TypeUser
// sourceData() map[string]string
func CheckBox(name string, data sourceData) *check {
	new := &check{
		name: name,
		Data: data,
	}

	return new
}

// check Box id y valor
type check struct {
	name                  string
	Data                  sourceData
	only_internal_contend bool
}

func (c check) InputName() string {
	return c.name
}

func (check) HtmlName() string {
	return "checkbox"
}

// validación con datos de entrada
func (c check) ValidateField(data_in string, skip_validation bool, options ...string) error {
	if !skip_validation {
		dataInArray := String().Split(data_in, ",")
		for _, idkeyIn := range dataInArray {
			if _, exists := (c.Data.SourceData())[idkeyIn]; !exists {
				if idkeyIn != "" {
					return errors.New("valor " + idkeyIn + " no corresponde al checkbox")
				} else {
					return errors.New("selección requerida")
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
