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

func (info) InputName() string {
	return "Info"
}

func (i info) HtmlName() string {
	return "text"
}

func (i info) BuildContainerView(id, field_name string, allow_skip_completed bool) string {
	return i.Value
}
