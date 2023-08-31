package input

import "github.com/cdvelop/model"

// SourceData() map[string]string
func DataList(data sourceData) model.Input {
	in := datalist{
		Data: data,
	}

	return model.Input{
		InputName: in.Name(),
		Tag:       &in,
		Validate:  &in,
		TestData:  &in,
	}
}

type datalist struct {
	Data sourceData
}

func (d datalist) Name() string {
	return d.HtmlName()
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
