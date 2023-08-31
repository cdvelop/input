package input

import "github.com/cdvelop/model"

// options: title="xxx".
//gender return {"f": "Femenino", "m": "Masculino"}.
// SourceData() map[string]string default: {"1": "Opción 1", "2": "Opción 2"}
func Radio(data sourceData, options ...string) model.Input {
	in := radio{
		name: "radio",
		Data: data,
		attributes: attributes{
			Onchange: `onchange="RadioChange(this);"`,
		},
	}
	in.Set(options...)

	for _, opt := range options {
		switch opt {
		case "gender":
			in.name = "gender"
			in.Data = gender{}

		}
	}

	if in.Data == nil {
		in.Data = radioDefault{}
	}

	return model.Input{
		InputName: in.Name(),
		Tag:       &in,
		Validate:  &in,
		TestData:  &in,
	}
}

type radio struct {
	name string
	Data sourceData
	attributes
}

func (r radio) Name() string {
	return r.name
}

func (radio) HtmlName() string {
	return "radio"
}

// validación con datos de entrada
func (r radio) ValidateField(data_in string, skip_validation bool, options ...string) error {
	if !skip_validation {
		if _, exists := r.Data.SourceData()[data_in]; !exists {

			if data_in != "" {
				return model.Error("valor", data_in, "no corresponde a botón radio")

			} else {
				return model.Error("selección requerida")
			}
		}
	}
	return nil
}

func (r radio) GoodTestData() (out []string) {
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
