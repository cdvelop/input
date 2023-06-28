package input

import "github.com/cdvelop/model"

func RadioGenero() model.Input {
	return Radio(genero{})
}

type genero struct{}

func (genero) SourceData() map[string]string {
	return map[string]string{"f": "Femenino", "m": "Masculino"}
}
