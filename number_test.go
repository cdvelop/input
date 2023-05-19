package input_test

import (
	"log"
	"testing"

	"github.com/cdvelop/input"
)

var (
	modelNumber = input.Number("", "")

	dataNumber = map[string]struct {
		inputData       string
		skip_validation bool
		expected        bool
	}{
		"numero correcto 100": {"100", false, true},
		"un carácter 0":       {"0", false, true},
		"un carácter 1":       {"1", false, true},
		"uint64 +20 char":     {"18446744073709551615", false, true},
		"uint64 +21 char?":    {"184467440737095516150", false, false},
		"int 64 -o+ 19 char":  {"9223372036854775807", false, true},
		"int 32 -o+ 10  char": {"2147483647", false, true},
		"18 cifras":           {"100002323262637278", false, true},

		"grande y letra 10000232e26263727": {"10000232E26263727", false, false},
		"numero negativo -100 ":            {"-100", false, false},
		"texto en vez de numero ":          {"h", false, false},
		"texto y numero":                   {"h1", false, false},
	}
)

func Test_InputNumber(t *testing.T) {
	for prueba, data := range dataNumber {
		t.Run((prueba), func(t *testing.T) {
			if ok := modelNumber.Validate.ValidateField(data.inputData, data.skip_validation); ok != data.expected {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

var (
	// 1 408 XXX XXXX
	// 5 699 524 9966

	modelPhoneNumber = input.Number(`pattern="^[0-9]{7,11}$"`)

	dataPhoneNumber = map[string]struct {
		inputData       string
		skip_validation bool
		expected        bool
	}{
		"numero correcto 7 dígitos":      {"1234567", false, true},
		"numero correcto 9 dígitos":      {"123456789", false, true},
		"numero correcto 11 dígitos ok?": {"12345678911", false, true},
		"con código país":                {"56988765432", false, true},
		"signo mas + ok?":                {"+56988765432", false, false},
		"numero correcto 6 dígitos ok?":  {"123456", false, false},
		"numero correcto 1 dígitos ok?":  {"0", false, false},
	}
)

func Test_InputPhoneNumber(t *testing.T) {
	for prueba, data := range dataPhoneNumber {
		t.Run((prueba), func(t *testing.T) {
			if ok := modelPhoneNumber.Validate.ValidateField(data.inputData, data.skip_validation); ok != data.expected {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_GoodInputPhoneNumber(t *testing.T) {
	for _, data := range modelPhoneNumber.TestData.GoodTestData("", "", true) {
		t.Run((data), func(t *testing.T) {
			if ok := modelPhoneNumber.Validate.ValidateField(data, false); !ok {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_GoodInputNumber(t *testing.T) {
	for _, data := range modelNumber.TestData.GoodTestData("", "", true) {
		t.Run((data), func(t *testing.T) {
			if ok := modelNumber.Validate.ValidateField(data, false); !ok {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_WrongInputNumber(t *testing.T) {
	for _, data := range modelNumber.TestData.WrongTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelNumber.Validate.ValidateField(data, false); ok {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}
