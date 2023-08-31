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
		expected        string
	}{
		"todo los caracteres permitidos?":   {"hola: esto, es. la - prueba 10", false, ""},
		"salto de linea permitido? y guion": {"hola:\n esto, es. la - \nprueba 10", false, ""},
		"letra ñ permitida? paréntesis y $": {"soy ñato o Ñato (aqui) costo $10000.", false, ""},
		"solo texto y espacio?":             {"hola esto es una prueba", false, ""},
		"texto y puntuación?":               {"hola: esto es una prueba", false, ""},
		"texto y puntuación y coma?":        {"hola: esto,true, es una prueba", false, ""},
		"4 caracteres?":                     {" .s5", false, ""},
		"sin data permitido?":               {"", false, "tamaño mínimo 2 caracteres"},
		"# permitido?":                      {"# son", false, ""},
		"¿ ? permitido?":                    {" ¿ si ?", false, "carácter ¿ no permitido"},
		"tildes si?":                        {" mí tílde", false, ""},
		"1 carácter?":                       {"1", false, "tamaño mínimo 2 caracteres"},
		"nombre correcto?":                  {"Dr. Pato Gomez", false, ""},
		"solo espacio en blanco?":           {" ", false, "tamaño mínimo 2 caracteres"},
		"texto largo correcto?":             {`IRRITACION EN PIEL DE ROSTRO. ALERGIAS NO. CIRUGIAS NO. ACTUAL TTO CON ISOTRETINOINA 10MG - ENERO 2022. EN TTO ACTUAL CON VIT D. EXAMEN DE LAB 20-12-2022. SIN OTROS ANTECEDENTES`, false, ""},
	}
)

func Test_TagTextArea(t *testing.T) {
	tag := modelTextArea.Tag.HtmlTag("1", "name", true)
	if tag == "" {
		log.Fatalln("ERROR NO TAG RENDERING ")
	}
}

func Test_InputTextArea(t *testing.T) {
	for prueba, data := range dataTextArea {
		t.Run((prueba + data.inputData), func(t *testing.T) {
			err := modelTextArea.Validate.ValidateField(data.inputData, data.skip_validation)
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
func Test_GoodInputTextArea(t *testing.T) {
	for _, data := range modelTextArea.TestData.GoodTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelTextArea.Validate.ValidateField(data, false); ok != nil {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_WrongInputTextArea(t *testing.T) {
	for _, data := range modelTextArea.TestData.WrongTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelTextArea.Validate.ValidateField(data, false); ok == nil {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}
