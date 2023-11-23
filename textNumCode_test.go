package input_test

import (
	"log"
	"testing"

	"github.com/cdvelop/input"
)

var (
	modelTextNumCode = input.TextNumCode()

	dataTextNumCode = map[string]struct {
		inputData       string
		skip_validation bool
		expected        string
	}{
		"2 letras un numero ":            {"et1", false, ""},
		"código venta ":                  {"V22400", false, ""},
		"código numero y letras":         {"12f", false, ""},
		"guion bajo permitido?":          {"son_24_botellas", false, ""},
		"espacio permitido?":             {"1os cuatro", false, "espacios en blanco no permitidos"},
		"palabras guion_bajo si? ":       {"son_2_cuadros", false, ""},
		"palabras separadas si?":         {"son 2 cuadros", false, "espacios en blanco no permitidos"},
		"palabras guion medio si?":       {"son-2-cuadros", false, ""},
		"solo texto ok":                  {"tres", false, ""},
		"friday ok":                      {"friday", false, ""},
		"saturday ok":                    {"saturday", false, ""},
		"wednesday ok":                   {"Wednesday", false, ""},
		"month 10 ok":                    {"10", false, ""},
		"month 03 ok":                    {"03", false, ""},
		"solo un carácter":               {"3", false, "tamaño mínimo 2 caracteres"},
		"guion al inicio ? 2 caracteres": {"-1", false, "no se puede comenzar con -"},
		"/ al inicio 2 caracteres":       {"/1", false, "no se puede comenzar con /"},
	}
)

func Test_InputTextNumCode(t *testing.T) {
	for prueba, data := range dataTextNumCode {
		t.Run((prueba + data.inputData), func(t *testing.T) {
			err := modelTextNumCode.Validate.ValidateField(data.inputData, data.skip_validation)

			if err != data.expected {
				log.Println(prueba)
				log.Fatalf("resultado: [%v] expectativa: [%v]\n%v", err, data.expected, data.inputData)
			}
		})
	}
}
func Test_TagTextNumCode(t *testing.T) {
	tag := modelTextNumCode.Tag.BuildContainerView("1", "name", true)
	if tag == "" {
		log.Fatalln("ERROR NO TAG RENDERING ")
	}
}

func Test_GoodInputTextNumCode(t *testing.T) {
	for _, data := range modelTextNumCode.TestData.GoodTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelTextNumCode.Validate.ValidateField(data, false); ok != "" {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_WrongInputTextNumCode(t *testing.T) {
	for _, data := range modelTextNumCode.TestData.WrongTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelTextNumCode.Validate.ValidateField(data, false); ok == "" {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}
