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

func (f filePath) InputName() string {
	return "FilePath"
}

func (f filePath) HtmlName() string {
	return "file"
}

func (f filePath) BuildContainerView(id, field_name string, allow_skip_completed bool) (tags string) {
	return f.BuildHtmlTag(f.HtmlName(), "FilePath", id, field_name, allow_skip_completed)
}

var errPath = errors.New("La ruta no puede comenzar con \\ o / ")

// validación con datos de entrada
func (f filePath) ValidateField(data_in string, skip_validation bool, options ...string) error {
	if !skip_validation {
		if data_in == "" {
			return errors.New("La ruta no puede estar vacía")
		}

		if data_in[0] == '\\' {
			return errPath
		}

		// Reemplazar las barras diagonales hacia adelante con barras diagonales hacia atrás.
		data_in = String().Replace(data_in, "/", "\\")

		// Eliminar barras diagonales dobles al principio y al final de la cadena.
		data_in = String().Replace(data_in, "\\", "")

		// Dividir la cadena en partes utilizando las barras diagonales como delimitadores.
		parts := String().Split(data_in, "\\")

		for _, part := range parts {
			err := f.per.Validate(part)
			if err != nil {
				return err
			}
		}

		// Verificar que la ruta sea válida para Linux y Windows
	}
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
