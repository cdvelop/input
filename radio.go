package input

import "errors"

// name ej: FileType,RadioGender...
// SourceData() map[string]string default: {"1": "Opción 1", "2": "Opción 2"}
func Radio(name string, data sourceData) *radio {
	new := &radio{
		name: name,
		Data: data,
		attributes: attributes{
			Onchange: `onchange="RadioChange(this);"`,
		},
	}

	if new.Data == nil {
		new.Data = radioDefault{}
	}

	return new
}

type radio struct {
	name string
	Data sourceData
	attributes
}

func (r radio) InputName(customName, htmlName *string) {
	if customName != nil {
		*customName = r.name
	}
	if htmlName != nil {
		*htmlName = "radio"
	}
}

// validación con datos de entrada
func (r radio) ValidateInput(value string) error {
	if _, exists := r.Data.SourceData()[value]; !exists {
		if value != "" {
			return errors.New("valor " + value + " no corresponde a botón radio")
		} else {
			return errors.New("selección requerida")
		}
	}
	return nil
}

func (r radio) GoodTestData() (out []string) {
	for k := range r.Data.SourceData() {
		out = append(out, k)
	}
	return
}

func (r radio) WrongTestData() (out []string) {
	for _, wd := range wrong_data {
		if _, exist := r.Data.SourceData()[wd]; !exist {
			out = append(out, wd)
		}
	}
	return
}
