package input_test

import (
	"log"
	"testing"

	"github.com/cdvelop/input"
)

var (
	modelPassword = input.Password()

	dataPassword = map[string]struct {
		inputData       string
		skip_validation bool
		expected        bool
	}{
		"valida numero letras y carácter": {"c0ntra3", false, true},
		"valida muchos caracteres":        {"M1 contraseÑ4", false, true},
		"valida 8 caracteres":             {"contrase", false, true},
		"valida 5 caracteres":             {"contñ", false, true},
		"valida solo números":             {"12345", false, true},
		"no valida menos de 2":            {"1", false, false},
		"sin data":                        {"", false, false},
	}
)

func Test_InputPassword(t *testing.T) {
	for prueba, data := range dataPassword {
		t.Run((prueba + ": " + data.inputData), func(t *testing.T) {
			if ok := modelPassword.Validate.ValidateField(data.inputData, data.skip_validation); ok != data.expected {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_GoodInputPassword(t *testing.T) {
	for _, data := range modelPassword.TestData.GoodTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelPassword.Validate.ValidateField(data, false); !ok {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_WrongInputPassword(t *testing.T) {
	for _, data := range modelPassword.TestData.WrongTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelPassword.Validate.ValidateField(data, false); ok {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

var modelPasswordMinimal = input.Password(`min="10"`, `max="30"`)

func Test_GoodInputPasswordMinimal(t *testing.T) {
	for _, data := range modelPasswordMinimal.TestData.GoodTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelPasswordMinimal.Validate.ValidateField(data, false); !ok {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_WrongInputPasswordMinimal(t *testing.T) {
	for _, data := range modelPasswordMinimal.TestData.WrongTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelPasswordMinimal.Validate.ValidateField(data, false); ok {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}
