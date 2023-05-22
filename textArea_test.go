package input_test

import (
	"log"
	"testing"

	"github.com/cdvelop/input"
)

var (
	modelTextArea = input.TextArea()

	dataTextArea = map[string]struct {
		inputData       string
		skip_validation bool
		expected        bool
	}{
		"todo los caracteres permitidos?":   {"hola: esto, es. la - prueba 10", false, true},
		"salto de linea permitido? y guion": {"hola:\n esto, es. la - \nprueba 10", false, true},
		"letra ñ permitida? paréntesis y $": {"soy ñato o Ñato (aqui) costo $10000.", false, true},
		"solo texto y espacio?":             {"hola esto es una prueba", false, true},
		"texto y puntuación?":               {"hola: esto es una prueba", false, true},
		"texto y puntuación y coma?":        {"hola: esto,true, es una prueba", false, true},
		"4 caracteres?":                     {" .s5", false, true},
		"sin data permitido?":               {"", false, false},
		"# permitido?":                      {"# son", false, false},
		"¿ ? permitido?":                    {" ¿ si ?", false, false},
		"tildes si?":                        {" mí tílde", false, true},
		"1 carácter?":                       {"1", false, false},
		"nombre correcto?":                  {"Dr. Pato Gome", false, true},
		"solo espacio en blanco?":           {" ", false, false},
		"texto largo correcto?":             {`IRRITACION EN PIEL DE ROSTRO. ALERGIAS NO. CIRUGIAS NO. ACTUAL TTO CON ISOTRETINOINA 10MG - ENERO 2022. EN TTO ACTUAL CON VIT D. EXAMEN DE LAB 20-12-2022. SIN OTROS ANTECEDENTES`, false, true},
	}
)

func Test_InputTextArea(t *testing.T) {
	for prueba, data := range dataTextArea {
		t.Run((prueba + data.inputData), func(t *testing.T) {
			if ok := modelTextArea.Validate.ValidateField(data.inputData, data.skip_validation); ok != data.expected {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}
func Test_GoodInputTextArea(t *testing.T) {
	for _, data := range modelTextArea.TestData.GoodTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelTextArea.Validate.ValidateField(data, false); !ok {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_WrongInputTextArea(t *testing.T) {
	for _, data := range modelTextArea.TestData.WrongTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelTextArea.Validate.ValidateField(data, false); ok {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}
