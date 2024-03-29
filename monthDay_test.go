package input_test

import (
	"log"
	"testing"

	"github.com/cdvelop/input"
)

var (
	modelMonthDay = input.MonthDay()

	dataMonthDay = map[string]struct {
		inputData       string
		skip_validation bool
		expected        string
	}{
		"dia ok?":                     {"01", false, ""},
		"caracteres?":                 {"0l", false, "l no es un numero"},
		"mes ok?":                     {"31", false, ""},
		"fecha recortada sin año ok?": {"31-01", false, "tamaño máximo 2 caracteres"},
		"correcto ?":                  {"1-1", false, "tamaño máximo 2 caracteres"},
		"incorrecto ":                 {"2002-12-03", false, "tamaño máximo 2 caracteres"},
		"formato incorrecto ":         {"21/12", false, "tamaño máximo 2 caracteres"},
		"data incorrecta ":            {"0000-00-00", false, "tamaño máximo 2 caracteres"},
		"toda la data correcta?":      {"", false, "tamaño mínimo 2 caracteres"},
	}
)

func Test_InputMonthDay(t *testing.T) {
	for prueba, data := range dataMonthDay {
		t.Run((prueba + data.inputData), func(t *testing.T) {
			err := modelMonthDay.ValidateField(data.inputData, data.skip_validation)

			if err != data.expected {
				log.Println(prueba)
				log.Fatalf("resultado: [%v] expectativa: [%v]\n%v", err, data.expected, data.inputData)
			}
		})
	}
}

func Test_TagMonthDay(t *testing.T) {
	tag := modelMonthDay.BuildContainerView("1", "name", true)
	if tag == "" {
		log.Fatalln("ERROR NO TAG RENDERING ")
	}
}

func Test_GoodInputMonthDay(t *testing.T) {
	for _, data := range modelMonthDay.GoodTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelMonthDay.ValidateField(data, false); ok != "" {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_WrongInputMonthDay(t *testing.T) {
	for _, data := range modelMonthDay.WrongTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelMonthDay.ValidateField(data, false); ok == "" {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}
