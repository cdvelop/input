package input_test

import (
	"log"
	"testing"

	"github.com/cdvelop/input"
)

var (
	modelTexOnly = input.TextOnly()

	dataTextOnly = map[string]struct {
		inputData       string
		skip_validation bool
		expected        bool
	}{
		"nombre correcto con punto?":       {"Dr. Maria Jose Diaz Cadiz", false, false},
		"palabras con tilde?":              {"María Jose Diáz Cadíz", false, true},
		"tilde ok ? ":                      {"peréz del rozal", false, true},
		"texto con ñ?":                     {"Ñuñez perez", false, true},
		"texto correcto + 3 caracteres ":   {"juli", false, true},
		"texto correcto 3 caracteres ":     {"luz", false, true},
		"oración ok ":                      {"hola que tal", false, true},
		"Dato numérico 100 no permitido? ": {"100", false, false},
		"con caracteres y coma ?":          {"los,true, vengadores", false, false},
		"sin data ok":                      {"", false, false},
		"un carácter numérico ?":           {"8", false, false},
		"palabra mas numero permitido ?":   {"son 4 bidones", false, false},
		"con paréntesis y numero ?":        {"son 4 (4 bidones)", false, false},
		"con solo paréntesis ?":            {"son (bidones)", false, false},
		"palabras y numero ?":              {"apellido Actualizado 1", false, false},
		"caracteres 47 ok?":                {"juan marcos antonio del rosario de las carmenes", false, true},
		"un carácter ok?":                  {"!", false, false},
	}
)

func Test_InputTextOnly(t *testing.T) {
	for prueba, data := range dataTextOnly {
		t.Run((prueba + data.inputData), func(t *testing.T) {
			if ok := modelTexOnly.Validate.ValidateField(data.inputData, data.skip_validation); ok != data.expected {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_GoodInputTextOnly(t *testing.T) {
	for _, data := range modelTexOnly.TestData.GoodTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelTexOnly.Validate.ValidateField(data, false); !ok {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_WrongInputTextOnly(t *testing.T) {
	for _, data := range modelTexOnly.TestData.WrongTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelTexOnly.Validate.ValidateField(data, false); ok {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}
