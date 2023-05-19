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
		expected        bool
	}{
		"2 letras un numero ":      {"et1", false, true},
		"código venta ":            {"V22400", false, true},
		"código numero y letras":   {"12f", false, true},
		"guion bajo permitido?":    {"son_24_botellas", false, false},
		"espacio permitido?":       {"1os cuatro", false, false},
		"palabras guion bajo si? ": {"son_2_cuadros", false, false},
		"palabras separadas si?":   {"son 2 cuadros", false, false},
		"palabras guion medio si?": {"son-2-cuadros", false, true},
		"solo texto ok":            {"tres", false, true},
		"friday ok":                {"friday", false, true},
		"saturday ok":              {"saturday", false, true},
		"wednesday ok":             {"Wednesday", false, true},
		"month 10 ok":              {"10", false, true},
		"month 03 ok":              {"03", false, true},
		"solo un carácter":         {"3", false, false},
		"2 caracteres":             {"-1", false, false},
	}
)

func Test_InputTextNumCode(t *testing.T) {
	for prueba, data := range dataTextNumCode {
		t.Run((prueba + data.inputData), func(t *testing.T) {
			if ok := modelTextNumCode.Validate.ValidateField(data.inputData, data.skip_validation); ok != data.expected {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_GoodInputTextNumCode(t *testing.T) {
	for _, data := range modelTextNumCode.TestData.GoodTestData("", "", true) {
		t.Run((data), func(t *testing.T) {
			if ok := modelTextNumCode.Validate.ValidateField(data, false); !ok {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_WrongInputTextNumCode(t *testing.T) {
	for _, data := range modelTextNumCode.TestData.WrongTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelTextNumCode.Validate.ValidateField(data, false); ok {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}
