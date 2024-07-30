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
	name string
	Data sourceData
}

func (d datalist) InputName() string {
	return d.name
}

func (datalist) HtmlName() string {
	return "datalist"
}

//  validación con datos de entrada
func (d datalist) ValidateField(data_in string, skip_validation bool, options ...string) error {
	if !skip_validation {
		if _, exists := d.Data.SourceData()[data_in]; !exists {
			if data_in != "" {
				return errors.New("valor " + data_in + " no permitido en datalist")
			} else {
				return errors.New("selección requerida")
			}
		}
	}
	return nil
}
