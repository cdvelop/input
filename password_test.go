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
		expected        string
	}{
		"valida numero letras y carácter": {"c0ntra3", false, ""},
		"valida muchos caracteres":        {"M1 contraseÑ4", false, ""},
		"valida 8 caracteres":             {"contrase", false, ""},
		"valida 5 caracteres":             {"contñ", false, ""},
		"valida solo números":             {"12345", false, ""},
		"no valida menos de 2":            {"1", false, "tamaño mínimo 5 caracteres"},
		"sin data":                        {"", false, "tamaño mínimo 5 caracteres"},
	}
)

func Test_InputPassword(t *testing.T) {
	for prueba, data := range dataPassword {
		t.Run((prueba + ": " + data.inputData), func(t *testing.T) {
			err := modelPassword.Validate.ValidateField(data.inputData, data.skip_validation)
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

func Test_TagPassword(t *testing.T) {
	tag := modelPassword.Tag.BuildContainerView("1", "name", true)
	if tag == "" {
		log.Fatalln("ERROR NO TAG RENDERING ")
	}
}

func Test_GoodInputPassword(t *testing.T) {
	for _, data := range modelPassword.TestData.GoodTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelPassword.Validate.ValidateField(data, false); ok != nil {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_WrongInputPassword(t *testing.T) {
	for _, data := range modelPassword.TestData.WrongTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelPassword.Validate.ValidateField(data, false); ok == nil {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

var modelPasswordMinimal = input.Password(`min="10"`, `max="30"`)

func Test_GoodInputPasswordMinimal(t *testing.T) {
	for _, data := range modelPasswordMinimal.TestData.GoodTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelPasswordMinimal.Validate.ValidateField(data, false); ok != nil {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_WrongInputPasswordMinimal(t *testing.T) {
	for _, data := range modelPasswordMinimal.TestData.WrongTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelPasswordMinimal.Validate.ValidateField(data, false); ok == nil {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}
