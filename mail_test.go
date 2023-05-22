package input_test

import (
	"log"
	"testing"

	"github.com/cdvelop/input"
)

var (
	modelMail = input.Mail()

	dataMail = map[string]struct {
		inputData       string
		skip_validation bool
		expected        bool
	}{
		"correo normal ": {"mi.correo@mail.com", false, true},
		// "código numero y letras":   {"12f", false, true},
		// "guion bajo permitido?":    {"son_24_botellas", false, false},
		// "espacio permitido?":       {"1os cuatro", false, false},
		// "palabras guion bajo si? ": {"son_2_cuadros", false, false},
		// "palabras separadas si?":   {"son 2 cuadros", false, false},
		// "palabras guion medio si?": {"son-2-cuadros", false, true},
		// "solo texto ok":            {"tres", false, true},
		// "friday ok":                {"friday", false, true},
		// "saturday ok":              {"saturday", false, true},
		// "wednesday ok":             {"Wednesday", false, true},
		// "month 10 ok":              {"10", false, true},
		// "month 03 ok":              {"03", false, true},
		// "solo un carácter":         {"3", false, false},
		// "2 caracteres":             {"-1", false, false},
	}
)

func Test_InputMail(t *testing.T) {
	for prueba, data := range dataMail {
		t.Run((prueba + data.inputData), func(t *testing.T) {
			if ok := modelMail.Validate.ValidateField(data.inputData, data.skip_validation); ok != data.expected {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_GoodInputMail(t *testing.T) {
	for _, data := range modelMail.TestData.GoodTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelMail.Validate.ValidateField(data, false); !ok {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_WrongInputMail(t *testing.T) {
	for _, data := range modelMail.TestData.WrongTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelMail.Validate.ValidateField(data, false); ok {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}
