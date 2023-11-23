package input_test

import (
	"log"
	"testing"

	"github.com/cdvelop/input"
)

var (
	modelDNI = input.Rut("dni-mode")

	dataDNI = map[string]struct {
		inputData       string
		skip_validation bool
		expected        string
	}{
		"ok 17734478-8":               {"17734478-8", false, ""},
		"ok 7863697-1":                {"7863697-1", false, ""},
		"ok 20373221-K":               {"20373221-k", false, ""},
		"run validado W? permitido?":  {"7863697-W", false, "dígito verificador W inválido"},
		"cambio dígito a k 7863697-k": {"7863697-k", false, "dígito verificador k inválido"},
		"cambio dígito a 0 7863697-0": {"7863697-0", false, "dígito verificador 0 inválido"},
		"ok 14080717-6":               {"14080717-6", false, ""},
		"incorrecto 14080717-0":       {"14080717-0", false, "dígito verificador 0 inválido"},
		"correcto cero al inicio? ":   {"07863697-1", false, "primer dígito no puede ser 0"},
		"data correcta solo espacio?": {" ", false, "tamaño mínimo 9 caracteres"},
		"caracteres permitidos?":      {`%$"1 uut4%%oo`, false, "% no es un numero"},
		"pasaporte ax001223b ok?":     {"ax001223b", false, ""},
		"caída con dato":              {"123", false, "tamaño mínimo 9 caracteres"},
	}
)

func Test_TagDNI(t *testing.T) {
	tag := modelDNI.Tag.BuildContainerView("1", "name", true)
	if tag == "" {
		log.Fatalln("ERROR NO TAG RENDERING ")
	}
}

func Test_InputDNI(t *testing.T) {
	for prueba, data := range dataDNI {
		t.Run((prueba), func(t *testing.T) {
			err := modelDNI.Validate.ValidateField(data.inputData, data.skip_validation)

			if err != data.expected {
				log.Println(prueba)
				log.Fatalf("resultado: [%v] expectativa: [%v]\n%v", err, data.expected, data.inputData)
			}
		})
	}
}

func Test_GoodInputDNI(t *testing.T) {
	for _, data := range modelDNI.TestData.GoodTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelDNI.Validate.ValidateField(data, false); ok != "" {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_WrongInputDNI(t *testing.T) {
	for _, data := range modelDNI.TestData.WrongTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelDNI.Validate.ValidateField(data, false); ok == "" {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}
