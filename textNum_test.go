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
		expected        bool
	}{
		"guion bajo ":           {"son_24_botellas", false, true},
		"frase con guion bajo ": {"los_cuatro", false, true},
		"frase sin guion bajo ": {"los cuatro", false, false},
		"palabras guion bajo ":  {"son_2_cuadros", false, true},
		"palabras separadas ":   {"son 2 cuadros", false, false},
		"palabras guion medio ": {"son-2-cuadros", false, false},
		"menos de 5 palabras ":  {"tres", false, false},
		"2 letras un numero ":   {"et1_", false, false},
	}
)

func Test_InputTextNum(t *testing.T) {
	for prueba, data := range dataTextNum {
		t.Run((prueba + data.inputData), func(t *testing.T) {
			if ok := modelTextNum.Validate.ValidateField(data.inputData, data.skip_validation); ok != data.expected {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_GoodInputTextNum(t *testing.T) {
	for _, data := range modelTextNum.TestData.GoodTestData("", "", true) {
		t.Run((data), func(t *testing.T) {
			if ok := modelTextNum.Validate.ValidateField(data, false); !ok {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_WrongInputTextNum(t *testing.T) {
	for _, data := range modelTextNum.TestData.WrongTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelTextNum.Validate.ValidateField(data, false); ok {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}
