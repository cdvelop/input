package input

import "github.com/cdvelop/model"

// name ej: OptionUser
// SourceData() map[string]string
func SelecTag(name string, data model.SourceData) *model.Input {
	in := selecTag{
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

type selecTag struct {
	name string
	Data model.SourceData
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
