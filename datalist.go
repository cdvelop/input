package input

import "errors"

// name ej: InputOptions
// SourceData() map[string]string
func DataList(name string, data sourceData) *datalist {
	new := &datalist{
		name: name,
		Data: data,
	}

	return new
}

type datalist struct {
	name               string
	Data               sourceData
	AllowSkipCompleted bool
}

func (d datalist) InputName(customName, htmlName *string) {
	if customName != nil {
		*customName = d.name
	}
	if htmlName != nil {
		*htmlName = "datalist"
	}
}

//  validación con datos de entrada
func (d datalist) ValidateInput(value string) error {
	if _, exists := d.Data.SourceData()[value]; !exists {
		if value != "" {
			return errors.New("valor " + value + " no permitido en datalist")
		} else {
			return errors.New("selección requerida")
		}
	}
	return nil
}
