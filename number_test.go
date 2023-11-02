package input_test

import (
	"log"
	"testing"

	"github.com/cdvelop/input"
)

var (
	modelNumber = input.Number()

	dataNumber = map[string]struct {
		inputData       string
		skip_validation bool
		expected        string
	}{
		"numero correcto 100": {"100", false, ""},
		"un carácter 0":       {"0", false, ""},
		"un carácter 1":       {"1", false, ""},
		"uint64 +20 char":     {"18446744073709551615", false, ""},
		"uint64 +21 char?":    {"184467440737095516150", false, "tamaño máximo 20 caracteres"},
		"int 64 -o+ 19 char":  {"9223372036854775807", false, ""},
		"int 32 -o+ 10  char": {"2147483647", false, ""},
		"18 cifras":           {"100002323262637278", false, ""},

		"grande y letra":          {"10000232E26263727", false, "E no es un numero"},
		"numero negativo -100 ":   {"-100", false, "- no es un numero"},
		"texto en vez de numero ": {"h", false, "h no es un numero"},
		"texto y numero":          {"h1", false, "h no es un numero"},
	}
)

func Test_InputNumber(t *testing.T) {
	for prueba, data := range dataNumber {
		t.Run((prueba), func(t *testing.T) {
			err := modelNumber.Validate.ValidateField(data.inputData, data.skip_validation)
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

func Test_TagNumber(t *testing.T) {
	tag := modelNumber.Tag.BuildContainerView("1", "name", true)
	if tag == "" {
		log.Fatalln("ERROR NO TAG RENDERING ")
	}
}

var (
	// 1 408 XXX XXXX
	// 5 699 524 9966

	dataPhoneNumber = map[string]struct {
		inputData       string
		skip_validation bool
		expected        string
	}{
		"numero correcto 7 dígitos":      {"1234567", false, ""},
		"numero correcto 9 dígitos":      {"123456789", false, ""},
		"numero correcto 11 dígitos ok?": {"12345678911", false, ""},
		"con código país":                {"56988765432", false, ""},
		"signo mas + ok?":                {"+56988765432", false, "tamaño máximo 11 caracteres"},
		"numero correcto 6 dígitos ok?":  {"123456", false, "tamaño mínimo 7 caracteres"},
		"numero correcto 1 dígitos ok?":  {"0", false, "tamaño mínimo 7 caracteres"},
	}
)

func Test_InputPhoneNumber(t *testing.T) {
	for prueba, data := range dataPhoneNumber {
		t.Run((prueba), func(t *testing.T) {
			err := input.Phone().ValidateField(data.inputData, data.skip_validation)
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

func Test_GoodInputPhoneNumber(t *testing.T) {
	for _, data := range input.Phone().GoodTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := input.Phone().ValidateField(data, false); ok != nil {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_GoodInputNumber(t *testing.T) {
	for _, data := range modelNumber.TestData.GoodTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelNumber.Validate.ValidateField(data, false); ok != nil {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_WrongInputNumber(t *testing.T) {
	for _, data := range modelNumber.TestData.WrongTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelNumber.Validate.ValidateField(data, false); ok == nil {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}
