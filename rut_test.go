package input

import (
	"fmt"
	"log"
	"testing"
)

var (
	modelRut = Rut()

	dataRut = map[string]struct {
		inputData       string
		skip_validation bool
		result          bool
	}{
		"ok 17734478-8":               {"17734478-8", false, true},
		"ok 7863697-1":                {"7863697-1", false, true},
		"ok 20373221-K":               {"20373221-k", false, true},
		"run validado? permitido?":    {"7863697-W", false, false},
		"cambio dígito a k 7863697-k": {"7863697-k", false, false},
		"cambio dígito a 0 7863697-0": {"7863697-0", false, false},
		"ok 14080717-6":               {"14080717-6", false, true},
		"incorrecto 14080717-0":       {"14080717-0", false, false},
		"correcto cero al inicio? ":   {"07863697-1", false, false},
		"data correcta solo espacio?": {" ", false, false},
		"caracteres permitidos?":      {`%$"1 `, false, false},
	}
)

func Test_InputRut(t *testing.T) {
	for prueba, data := range dataRut {
		t.Run((prueba), func(t *testing.T) {
			if ok := modelRut.Validate.ValidateField(data.inputData, data.skip_validation); ok != data.result {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

var (
	dataRuNormalize = map[string]struct {
		inputData string
		expected  string
		result    bool
	}{
		"cero agregado al inicio?":         {"07863697-1", "7863697-1", true},
		"1 dígito verificador incorrecto?": {"7863697-0", "7863697-1", true},
		"1 cero agregado al inicio?":       {"05335341-K", "05335341-K", true},

		"dígito verificador incorrecto?": {"20373221-0", "20373221-K", true},

		"sin espacio en blanco?":     {" ", "", false},
		"nada de espacio en blanco?": {"    ", "", false},
	}
)

func Test_RutNormalize(t *testing.T) {
	for prueba, data := range dataRuNormalize {
		t.Run((prueba), func(t *testing.T) {
			if ok := RutNormalize(&data.inputData); ok != data.result {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
			// t.Logf("data entrada: [%v] resultado: [%v]\n", data.inputData, data.expected)
		})
	}
}

func Test_RutDigito(t *testing.T) {
	run := 17734478
	dv := DvRut(run)
	fmt.Printf("RUN: %v DIGITO: %v", run, dv)
}

func Test_GoodInputRut(t *testing.T) {
	for _, data := range modelRut.TestData.GoodTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelRut.Validate.ValidateField(data, false); !ok {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_WrongInputRut(t *testing.T) {
	for _, data := range modelRut.TestData.WrongTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelRut.Validate.ValidateField(data, false); ok {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}
