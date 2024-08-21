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
	name                string
	Data                sourceData
	onlyInternalContend bool
}

func (c check) InputName(customName, htmlName *string) {
	if customName != nil {
		*customName = c.name
	}
	if htmlName != nil {
		*htmlName = "checkbox"
	}
}

// validación con datos de entrada
func (c check) ValidateInput(value string) error {

	dataInArray := String().Split(value, ",")
	for _, idkeyIn := range dataInArray {
		if _, exists := (c.Data.SourceData())[idkeyIn]; !exists {
			if idkeyIn != "" {
				return errors.New("valor " + idkeyIn + " no corresponde al checkbox")
			} else {
				return errors.New("selección requerida")
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
