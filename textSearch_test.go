package input_test

import (
	"log"
	"testing"

	"github.com/cdvelop/input"
)

var (
	modelTextSearch = input.TextSearch()

	dataTextSearch = map[string]struct {
		inputData       string
		skip_validation bool
		expected        bool
	}{
		"palabra solo texto 15 caracteres?": {"Maria Jose Diaz", false, true},
		"texto con ñ ok?":                   {"Ñuñez perez", false, true},
		"tilde permitido?":                  {"peréz del rozal", false, false},
		"mas de 20 caracteres permitidos?":  {"hola son mas de 21 ca", false, false},
		"guion permitido":                   {"12038-0", false, true},
		"fecha correcta?":                   {"1990-07-21", false, true},
		"fecha incorrecta permitida?":       {"190-07-21", false, true},
	}
)

func Test_InputTextSearch(t *testing.T) {
	for prueba, data := range dataTextSearch {
		t.Run((prueba + data.inputData), func(t *testing.T) {
			if ok := modelTextSearch.Validate.ValidateField(data.inputData, data.skip_validation); ok != data.expected {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_GoodInputTextSearch(t *testing.T) {
	for _, data := range modelTextSearch.TestData.GoodTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelTextSearch.Validate.ValidateField(data, false); !ok {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_WrongInputTextSearch(t *testing.T) {
	for _, data := range modelTextSearch.TestData.WrongTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelTextSearch.Validate.ValidateField(data, false); ok {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}
