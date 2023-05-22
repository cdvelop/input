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
		result          bool
	}{
		"correcto ":                   {"2002-12-03", false, true},
		"carácter de mas incorrecto ": {"2002-12-03-", false, false},
		"formato incorrecto ":         {"21/12/1998", false, false},
		"fecha incorrecta ":           {"2020-31-01", false, false},
		"fecha recortada sin año ok?": {"31-01", false, false},
		"data incorrecta ":            {"0000-00-00", false, false},
		"toda la data correcta?":      {"", false, false},
	}
)

func Test_InputDate(t *testing.T) {
	for prueba, data := range dataDate {
		t.Run((prueba + data.inputData), func(t *testing.T) {
			if ok := modelDate.Validate.ValidateField(data.inputData, data.skip_validation); ok != data.result {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_GoodInputDate(t *testing.T) {
	for _, data := range modelDate.TestData.GoodTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelDate.Validate.ValidateField(data, false); !ok {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_WrongInputDate(t *testing.T) {
	for _, data := range modelDate.TestData.WrongTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelDate.Validate.ValidateField(data, false); ok {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}
