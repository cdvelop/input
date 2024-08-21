package input

//value = valor a mostrar
func Info(value string) *info {

	return &info{
		Value: value,
	}
}

// input de carácter informativo
type info struct {
	Value string //valor a mostrar
}

func (i info) InputName(customName, htmlName *string) {
	if customName != nil {
		*customName = "Info"
	}
	if htmlName != nil {
		*htmlName = "text"
	}
}

func (i info) BuildInputHtml(id, fieldName string) string {
	return i.Value
}
