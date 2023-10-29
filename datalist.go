package input

import "github.com/cdvelop/model"

// name ej: InputOptions
// SourceData() map[string]string
func DataList(name string, data model.SourceData) *model.Input {
	in := datalist{
		name: name,
		Data: data,
	}

	return &model.Input{
		InputName: name,
		Tag:       &in,
		Validate:  &in,
		TestData:  &in,
	}
}

type datalist struct {
	name string
	Data model.SourceData
}

func (datalist) HtmlName() string {
	return "datalist"
}

//  validación con datos de entrada
func (d datalist) ValidateField(data_in string, skip_validation bool, options ...string) error {
	if !skip_validation {
		if _, exists := d.Data.SourceData()[data_in]; !exists {

			if data_in != "" {
				return model.Error("valor", data_in, "no permitido en datalist")
			} else {
				return model.Error("selección requerida")

			}
		}
	}
	return nil
}
