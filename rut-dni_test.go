package input_test

import (
	"log"
	"testing"

	"github.com/cdvelop/input"
)

var (
	modelRunGlobal = input.Rut("dni-mode")
	// ("", "", false, false)

	dataRunGlobal = map[string]struct {
		inputData       string
		skip_validation bool
		expected        bool
	}{
		"ok 17734478-8":               {"17734478-8", false, true},
		"ok 7863697-1":                {"7863697-1", false, true},
		"ok 20373221-K":               {"20373221-k", false, true},
		"run validado W? permitido?":  {"7863697-W", false, false},
		"cambio dígito a k 7863697-k": {"7863697-k", false, false},
		"cambio dígito a 0 7863697-0": {"7863697-0", false, false},
		"ok 14080717-6":               {"14080717-6", false, true},
		"incorrecto 14080717-0":       {"14080717-0", false, false},
		"correcto cero al inicio? ":   {"07863697-1", false, false},
		"data correcta solo espacio?": {" ", false, false},
		"caracteres permitidos?":      {`%$"1 `, false, false},
		"pasaporte ax001223b ok?":     {"ax001223b", false, true},
		"caída con dato":              {"123", false, false},
	}
)

func Test_InputRunGlobal(t *testing.T) {
	for prueba, data := range dataRunGlobal {
		t.Run((prueba), func(t *testing.T) {
			if ok := modelRunGlobal.Validate.ValidateField(data.inputData, data.skip_validation); ok != data.expected {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_GoodInputRunGlobal(t *testing.T) {
	for _, data := range modelRunGlobal.TestData.GoodTestData("", "", true) {
		t.Run((data), func(t *testing.T) {
			if ok := modelRunGlobal.Validate.ValidateField(data, false); !ok {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_WrongInputRunGlobal(t *testing.T) {
	for _, data := range modelRunGlobal.TestData.WrongTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelRunGlobal.Validate.ValidateField(data, false); ok {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}
