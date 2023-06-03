package input

import "github.com/cdvelop/model"

// SourceData() map[string]string
func SelecTag(data sourceData) model.Input {
	in := selecTag{
		Data: data,
	}

	return model.Input{
		Component: model.Component{
			Name:        in.Name(),
			CssGlobal:   nil,
			CssPrivate:  nil,
			JsGlobal:    nil,
			JsPrivate:   nil,
			JsListeners: nil,
		},
		Tag:      in,
		Validate: in,
		TestData: in,
	}
}

type selecTag struct {
	Data sourceData
}

func (s selecTag) Name() string {
	return "selecTag"
}

func (s selecTag) HtmlName() string {
	return "select"
}

// validación con datos de entrada
func (s selecTag) ValidateField(data_in string, skip_validation bool) bool {
	if !skip_validation {
		if data_in != "" {
			if _, exists := s.Data.SourceData()[data_in]; exists {
				return true
			}
		}
	} else {
		return true
	}
	return false
}

func (s selecTag) GoodTestData() (out []string) {
	for k := range s.Data.SourceData() {
		out = append(out, k)
	}
	return
}

func (s selecTag) WrongTestData() (out []string) {
	for _, wd := range wrong_data {
		if _, exist := s.Data.SourceData()[wd]; !exist {
			out = append(out, wd)
		}
	}
	return
}
