package input_test

import (
	"log"
	"testing"

	"github.com/cdvelop/input"
)

var (
	modelHour = input.Hour()

	dataHour = map[string]struct {
		inputData       string
		skip_validation bool
		expected        string
	}{
		"correcto":    {"23:59", false, ""},
		"correcto 00": {"00:00", false, ""},
		"correcto 12": {"12:00", false, ""},

		"incorrecto 24":       {"24:00", false, "la hora 24 no existe"},
		"incorrecto sin data": {"", false, "tamaño mínimo 5 caracteres"},
		"incorrecto carácter": {"13-34", false, "carácter - no permitido"},
	}
)

func Test_InputHour(t *testing.T) {
	for prueba, data := range dataHour {
		t.Run((prueba + " " + data.inputData), func(t *testing.T) {
			err := modelHour.Validate.ValidateField(data.inputData, data.skip_validation)
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

func Test_TagHour(t *testing.T) {
	tag := modelHour.Tag.HtmlTag("1", "name", true)
	if tag == "" {
		log.Fatalln("ERROR NO TAG RENDERING ")
	}
}

func Test_GoodInputHour(t *testing.T) {
	for _, data := range modelHour.TestData.GoodTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelHour.Validate.ValidateField(data, false); ok != nil {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_WrongInputHour(t *testing.T) {
	for _, data := range modelHour.TestData.WrongTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelHour.Validate.ValidateField(data, false); ok == nil {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}
