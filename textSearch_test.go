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
	tag := modelTextSearch.BuildContainerView("1", "name", true)
	if tag == "" {
		log.Fatalln("ERROR NO TAG RENDERING ")
	}
}

func Test_InputTextSearch(t *testing.T) {
	for prueba, data := range dataTextSearch {
		t.Run((prueba + data.inputData), func(t *testing.T) {
			err := modelTextSearch.ValidateField(data.inputData, data.skip_validation)

			var err_str string
			if err != nil {
				err_str = err.Error()
			}

			if err_str != data.expected {
				log.Println(prueba)
				log.Fatalf("resultado: [%v] expectativa: [%v]\n%v", err, data.expected, data.inputData)
			}
		})
	}
}

func Test_GoodInputTextSearch(t *testing.T) {
	for _, data := range modelTextSearch.GoodTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelTextSearch.ValidateField(data, false); ok != nil {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_WrongInputTextSearch(t *testing.T) {
	for _, data := range modelTextSearch.WrongTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelTextSearch.ValidateField(data, false); ok == nil {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}
