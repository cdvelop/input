package input

import "github.com/cdvelop/model"

//{"d":"Dama","v":"Varón"}
// options: title="xxx"
func Radio(data sourceData, options ...string) model.Input {
	in := radio{
		Data:       data,
		attributes: attributes{},
	}
	in.Set(options...)

	return model.Input{
		Component: model.Component{
			Name:        in.Name(),
			CssGlobal:   nil,
			CssPrivate:  nil,
			JsGlobal:    nil,
			JsPrivate:   nil,
			JsListeners: nil,
		},
		Build:    in,
		Validate: in,
		TestData: in,
	}
}

type radio struct {
	Data sourceData
	attributes
}

func (r radio) Name() string {
	return r.HtmlName()
}

func (radio) HtmlName() string {
	return "radio"
}

// validación con datos de entrada
func (r radio) ValidateField(data_in string, skip_validation bool) bool {
	if !skip_validation {

		if _, exists := r.Data.SourceData()[data_in]; !exists {
			return false
		}
		return true

	} else {
		return true
	}
	return false
}

func (r radio) GoodTestData(table_name, field_name string, random bool) (out []string) {
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
