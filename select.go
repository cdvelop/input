package input

import "errors"

// name ej: OptionUser
// SourceData() map[string]string
func SelecTag(name string, data sourceData) *selecTag {

	return &selecTag{
		name: name,
		Data: data,
	}
}

type selecTag struct {
	name string
	Data sourceData
}

func (s selecTag) InputName() string {
	return s.name
}

func (s selecTag) HtmlName() string {
	return "select"
}

// validación con datos de entrada
func (s selecTag) ValidateField(data_in string, skip_validation bool, options ...string) error {
	if !skip_validation {
		if _, exists := s.Data.SourceData()[data_in]; !exists {
			if data_in != "" {
				return errors.New("valor " + data_in + " no corresponde al select")
			} else {
				return errors.New("selección requerida")
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
