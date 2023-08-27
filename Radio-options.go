package input

type gender struct{}

func (gender) SourceData() map[string]string {
	return map[string]string{"f": "Femenino", "m": "Masculino"}
}

type radioDefault struct{}

func (radioDefault) SourceData() map[string]string {
	return map[string]string{"1": "Opción 1", "2": "Opción 2"}
}
