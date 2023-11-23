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
		expected        string
	}{
		"palabra solo texto 15 caracteres?": {"Maria Jose Diaz", false, ""},
		"texto con ñ ok?":                   {"Ñuñez perez", false, ""},
		"tilde permitido?":                  {"peréz del rozal", false, "é con tilde no permitida"},
		"mas de 20 caracteres permitidos?":  {"hola son mas de 21 ca", false, "tamaño máximo 20 caracteres"},
		"guion permitido":                   {"12038-0", false, ""},
		"fecha correcta?":                   {"1990-07-21", false, ""},
		"fecha incorrecta permitida?":       {"190-07-21", false, ""},
	}
)

func Test_TagTextSearch(t *testing.T) {
	tag := modelTextSearch.Tag.BuildContainerView("1", "name", true)
	if tag == "" {
		log.Fatalln("ERROR NO TAG RENDERING ")
	}
}

func Test_InputTextSearch(t *testing.T) {
	for prueba, data := range dataTextSearch {
		t.Run((prueba + data.inputData), func(t *testing.T) {
			err := modelTextSearch.Validate.ValidateField(data.inputData, data.skip_validation)

			if err != data.expected {
				log.Println(prueba)
				log.Fatalf("resultado: [%v] expectativa: [%v]\n%v", err, data.expected, data.inputData)
			}
		})
	}
}

func Test_GoodInputTextSearch(t *testing.T) {
	for _, data := range modelTextSearch.TestData.GoodTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelTextSearch.Validate.ValidateField(data, false); ok != "" {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_WrongInputTextSearch(t *testing.T) {
	for _, data := range modelTextSearch.TestData.WrongTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelTextSearch.Validate.ValidateField(data, false); ok == "" {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}
