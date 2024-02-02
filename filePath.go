package input

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

const errPath = "La ruta no puede comenzar con \\ o / "

// validación con datos de entrada
func (f filePath) ValidateField(data_in string, skip_validation bool, options ...string) (err string) {
	if !skip_validation {
		if data_in == "" {
			return "La ruta no puede estar vacía"
		}

		if data_in[0] == '\\' {
			return errPath
		}

		// if data_in == ".\\" { // ruta valida
		// 	return ""
		// }

		// Reemplazar las barras diagonales hacia adelante con barras diagonales hacia atrás.
		data_in = String().Replace(data_in, "/", "\\")

		// fmt.Println("ENTRADA: ", data_in)

		// Eliminar barras diagonales dobles al principio y al final de la cadena.
		data_in = String().Replace(data_in, "\\", "")
		// data_in = strings.Trim(data_in, "\\")

		// Dividir la cadena en partes utilizando las barras diagonales como delimitadores.
		parts := String().Split(data_in, "\\")

		// fmt.Println("PARTES: ", parts)

		for _, part := range parts {
			err := f.per.Validate(part)
			if err != "" {
				return err
			}
		}

		// Verificar que la ruta sea válida para Linux y Windows
	}
	return ""
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
