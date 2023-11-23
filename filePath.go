package input

import (
	"strings"

	"github.com/cdvelop/model"
)

// options:
// "multiple"
// accept="image/*"
// title="Imágenes jpg"
func FilePath(options ...string) *model.Input {
	in := filePath{
		attributes: attributes{
			// Pattern: `^(\.\/|\.\?\\|\/)?([\w\s.-]+[\\\/]?)*$`,
		},
		per: Permitted{
			Letters:    true,
			Tilde:      false,
			Numbers:    true,
			Characters: []rune{'\\', '/', '.'},
			Minimum:    1,
			Maximum:    100,
		},
	}
	in.Set(options...)

	return &model.Input{
		InputName: "FilePath",
		Tag:       &in,
		Validate:  &in,
		TestData:  &in,
	}
}

type filePath struct {
	attributes
	per Permitted
}

func (f filePath) HtmlName() string {
	return "file"
}

func (f filePath) BuildContainerView(id, field_name string, allow_skip_completed bool) (tags string) {
	return f.BuildHtmlTag(f.HtmlName(), "FilePath", id, field_name, allow_skip_completed)
}

// validación con datos de entrada
func (f filePath) ValidateField(data_in string, skip_validation bool, options ...string) (err string) {
	if !skip_validation {
		if data_in == "" {
			return "La ruta no puede estar vacía"
		}

		if data_in[0] == '\\' {
			return "La ruta no puede comenzar con \\ o / "
		}

		// Reemplazar las barras diagonales hacia adelante con barras diagonales hacia atrás.
		data_in = strings.ReplaceAll(data_in, "/", "\\")

		// fmt.Println("ENTRADA: ", data_in)

		// Eliminar barras diagonales dobles al principio y al final de la cadena.
		data_in = strings.Trim(data_in, "\\")

		// Dividir la cadena en partes utilizando las barras diagonales como delimitadores.
		parts := strings.Split(data_in, "\\")

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
