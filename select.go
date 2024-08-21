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
	name               string
	Data               sourceData
	AllowSkipCompleted bool
}

func (s selecTag) InputName(customName, htmlName *string) {
	if customName != nil {
		*customName = s.name
	}
	if htmlName != nil {
		*htmlName = "select"
	}
}

// validación con datos de entrada
func (s selecTag) ValidateInput(value string) error {
	if _, exists := s.Data.SourceData()[value]; !exists {
		if value != "" {
			return errors.New("valor " + value + " no corresponde al select")
		} else {
			return errors.New("selección requerida")
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
