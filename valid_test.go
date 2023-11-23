package input_test

import (
	"log"
	"testing"

	"github.com/cdvelop/input"
)

var (
	validTestData = map[string]struct {
		text     string
		expected string
		input.Permitted
	}{
		"números sin espacio ok":                    {"5648", "", input.Permitted{Numbers: true}},
		"números con espacio ok":                    {"5648 78212", "", input.Permitted{Numbers: true, Characters: []rune{' '}}},
		"error no permitido números con espacio":    {"5648 78212", "espacios en blanco no permitidos", input.Permitted{Numbers: true}},
		"solo texto sin espacio ok":                 {"Maria", "", input.Permitted{Letters: true}},
		"texto con espacios ok":                     {"Maria De Lourdes", "", input.Permitted{Letters: true, Characters: []rune{' '}}},
		"texto con tildes y espacios ok":            {"María Dé Lourdes", "", input.Permitted{Tilde: true, Letters: true, Characters: []rune{' '}}},
		"texto con numero sin espacios ok":          {"equipo01", "", input.Permitted{Letters: true, Numbers: true}},
		"numero al inicio y texto sin espacios ok":  {"9equipo01", "", input.Permitted{Letters: true, Numbers: true}},
		"numero al inicio y texto con espacios ok":  {"9equipo01 2equipo2", "", input.Permitted{Letters: true, Numbers: true, Characters: []rune{' '}}},
		"error solo números no letras si espacios ": {"9equipo01 2equipo2", "carácter e no permitido", input.Permitted{Numbers: true, Characters: []rune{' '}}},
		"correo con punto y @ ok":                   {"mi.correo1@mail.com", "", input.Permitted{Characters: []rune{'@', '.'}, Numbers: true, Letters: true}},
		"error correo con tilde no permitido":       {"mí.correo@mail.com", "í con tilde no permitida", input.Permitted{Characters: []rune{'@', '.'}, Numbers: true, Letters: true}},
	}
)

func Test_Valid(t *testing.T) {

	for prueba, data := range validTestData {
		t.Run((prueba + " " + data.text), func(t *testing.T) {
			err := data.Validate(data.text)

			if err != data.expected {
				log.Println(prueba)
				log.Fatalf("expectativa [%v] resultado [%v]\n%v", data.expected, err, data.text)
			}

		})
	}
}
