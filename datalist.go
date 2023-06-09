package input

import "github.com/cdvelop/model"

// SourceData() map[string]string
func DataList(data sourceData) model.Input {
	in := datalist{
		Data: data,
	}

	return model.Input{
		Object: model.Object{
			ApiHandler: model.ApiHandler{
				Name: in.Name(),
			},
			Css:         nil,
			JsGlobal:    nil,
			JsFunctions: nil,
			JsListeners: nil,
		},
		Tag:      in,
		Validate: in,
		TestData: in,
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
func (d datalist) ValidateField(data_in string, skip_validation bool) bool {
	if !skip_validation {
		if data_in != "" {
			if _, exists := d.Data.SourceData()[data_in]; exists {
				return true
			}
		}
	} else {
		return true
	}
	return false
}
