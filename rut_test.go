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
		expected        string
	}{
		"sin guion 15890022k":         {"15890022k", false, "rut incorrecto"},
		"no tiene guion 177344788":    {"177344788", false, "rut incorrecto"},
		"ok 7863697-1":                {"7863697-1", false, ""},
		"ok 20373221-K":               {"20373221-k", false, ""},
		"run validado? permitido?":    {"7863697-W", false, "dígito verificador W inválido"},
		"cambio dígito a k 7863697-k": {"7863697-k", false, "dígito verificador k inválido"},
		"cambio dígito a 0 7863697-0": {"7863697-0", false, "dígito verificador 0 inválido"},
		"ok 14080717-6":               {"14080717-6", false, ""},
		"incorrecto 14080717-0":       {"14080717-0", false, "dígito verificador 0 inválido"},
		"correcto cero al inicio? ":   {"07863697-1", false, "primer dígito no puede ser 0"},
		"data correcta solo espacio?": {" ", false, "rut sin información"},
		"ok 17734478-8":               {"17734478-8", false, ""},
		"caracteres permitidos?":      {`%$"1 `, false, "rut incorrecto"},
		"no tiene guion 20373221K":    {"20373221k", false, "rut incorrecto"},
	}
)

func Test_InputRut(t *testing.T) {
	for prueba, data := range dataRut {
		t.Run((prueba), func(t *testing.T) {
			err := modelRut.Validate.ValidateField(data.inputData, data.skip_validation)
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

func Test_TagRut(t *testing.T) {
	tag := modelRut.Tag.BuildContainerView("1", "name", true)
	if tag == "" {
		log.Fatalln("ERROR NO TAG RENDERING ")
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
			if ok := modelRut.Validate.ValidateField(data, false); ok != nil {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_WrongInputRut(t *testing.T) {
	for _, data := range modelRut.TestData.WrongTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelRut.Validate.ValidateField(data, false); ok == nil {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}
