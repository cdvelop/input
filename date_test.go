package input_test

import (
	"log"
	"testing"

	"github.com/cdvelop/input"
)

var (
	modelDate = input.Date()

	dataDate = map[string]struct {
		inputData       string
		skip_validation bool
		expected        string
	}{
		"correcto ":                   {"2002-12-03", false, ""},
		"carácter de mas incorrecto ": {"2002-12-03-", false, "formato fecha no válido"},
		"formato incorrecto ":         {"21/12/1998", false, "formato fecha no válido"},
		"fecha incorrecta ":           {"2020-31-01", false, "Mes no válido"},
		"fecha recortada sin año ok?": {"31-01", false, "formato fecha no válido"},
		"data incorrecta ":            {"0000-00-00", false, "fecha ejemplo no válida"},
		"toda la data correcta?":      {"", false, "formato fecha no válido"},
	}
)

func Test_InputDate(t *testing.T) {
	for prueba, data := range dataDate {
		t.Run((prueba + data.inputData), func(t *testing.T) {
			err := modelDate.Validate.ValidateField(data.inputData, data.skip_validation)
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

func Test_TagDate(t *testing.T) {
	tag := modelDate.Tag.HtmlTag("1", "name", true)
	if tag == "" {
		log.Fatalln("ERROR NO TAG RENDERING ")
	}
}

func Test_GoodInputDate(t *testing.T) {
	for _, data := range modelDate.TestData.GoodTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelDate.Validate.ValidateField(data, false); ok != nil {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_WrongInputDate(t *testing.T) {
	for _, data := range modelDate.TestData.WrongTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelDate.Validate.ValidateField(data, false); ok == nil {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}
