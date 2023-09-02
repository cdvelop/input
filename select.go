package input

import "github.com/cdvelop/model"

// SourceData() map[string]string
func SelecTag(data model.SourceData) model.Input {
	in := selecTag{
		Data: data,
	}

	return model.Input{
		InputName: in.Name(),
		Tag:       &in,
		Validate:  &in,
		TestData:  &in,
	}
}

type selecTag struct {
	Data model.SourceData
}

func (s selecTag) Name() string {
	return "selecTag"
}

func (s selecTag) HtmlName() string {
	return "select"
}

// validación con datos de entrada
func (s selecTag) ValidateField(data_in string, skip_validation bool, options ...string) error {
	if !skip_validation {
		if _, exists := s.Data.SourceData()[data_in]; !exists {
			if data_in != "" {
				return model.Error("valor", data_in, "no corresponde al select")
			} else {
				return model.Error("selección requerida")
			}
		}
	}
	return nil
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
