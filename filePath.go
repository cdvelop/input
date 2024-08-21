package input

import "errors"

// options:
// "multiple"
// accept="image/*"
// title="Imágenes jpg"
func FilePath(options ...string) *filePath {
	new := &filePath{
		attributes: attributes{
			// Pattern: `^(\.\/|\.\?\\|\/)?([\w\s.-]+[\\\/]?)*$`,
		},
		per: Permitted{
			Letters:    true,
			Tilde:      false,
			Numbers:    true,
			Characters: []rune{'\\', '/', '.', '_'},
			Minimum:    1,
			Maximum:    100,
		},
	}
	new.Set(options...)

	return new
}

type filePath struct {
	attributes
	per Permitted
}

func (f filePath) InputName(customName, htmlName *string) {
	if customName != nil {
		*customName = "FilePath"
	}
	if htmlName != nil {
		*htmlName = "file"
	}
}

func (f filePath) BuildInputHtml(id, fieldName string) (tags string) {
	return f.BuildHtmlTag("file", "FilePath", id, fieldName)
}

var errPath = errors.New("la ruta no puede comenzar con \\ o / ")

// validación con datos de entrada
func (f filePath) ValidateInput(value string) error {
	if value == "" {
		return errors.New("la ruta no puede estar vacía")
	}

	if value[0] == '\\' {
		return errPath
	}

	// Reemplazar las barras diagonales hacia adelante con barras diagonales hacia atrás.
	value = String().Replace(value, "/", "\\")

	// Eliminar barras diagonales dobles al principio y al final de la cadena.
	value = String().Replace(value, "\\", "")

	// Dividir la cadena en partes utilizando las barras diagonales como delimitadores.
	parts := String().Split(value, "\\")

	for _, part := range parts {
		err := f.per.Validate(part)
		if err != nil {
			return err
		}
	}

	// Verificar que la ruta sea válida para Linux y Windows
	return nil
}

func (f filePath) GoodTestData() (out []string) {

	return []string{
		".\\misArchivos",
		".\\todos\\videos",
	}
}

func (f filePath) WrongTestData() (out []string) {
	out = []string{
		"\\-",
		"///.",
	}
	return
}
