package input_test

import (
	"log"
	"testing"

	"github.com/cdvelop/input"
)

var (
	modelTextNum = input.TextNum()

	dataTextNum = map[string]struct {
		inputData       string
		skip_validation bool
		expected        string
	}{
		"guion bajo ":           {"son_24_botellas", false, ""},
		"frase con guion bajo ": {"los_cuatro", false, ""},
		"frase sin guion bajo ": {"los cuatro", false, "espacios en blanco no permitidos"},
		"palabras guion bajo ":  {"son_2_cuadros", false, ""},
		"palabras separadas ":   {"son 2 cuadros", false, "espacios en blanco no permitidos"},
		"palabras guion medio ": {"son-2-cuadros", false, "carácter - no permitido"},
		"menos de 5 palabras ":  {"tres", false, "tamaño mínimo 5 caracteres"},
		"2 letras un numero ":   {"et1_", false, "tamaño mínimo 5 caracteres"},
	}
)

func Test_TagTextNum(t *testing.T) {
	tag := modelTextNum.Tag.BuildContainerView("1", "name", true)
	if tag == "" {
		log.Fatalln("ERROR NO TAG RENDERING ")
	}
}

func Test_InputTextNum(t *testing.T) {
	for prueba, data := range dataTextNum {
		t.Run((prueba + data.inputData), func(t *testing.T) {
			err := modelTextNum.Validate.ValidateField(data.inputData, data.skip_validation)
			var resp string
			if err != nil {
				resp = err.Error()
			}

			if resp != data.expected {
				log.Println(prueba)
				log.Fatalf("resultado: [%v] expectativa: [%v]\n%v", resp, data.expected, data.inputData)
			}
		})
	}
}

func Test_GoodInputTextNum(t *testing.T) {
	for _, data := range modelTextNum.TestData.GoodTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelTextNum.Validate.ValidateField(data, false); ok != nil {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_WrongInputTextNum(t *testing.T) {
	for _, data := range modelTextNum.TestData.WrongTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelTextNum.Validate.ValidateField(data, false); ok == nil {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}
