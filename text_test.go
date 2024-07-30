package input_test

import (
	"log"
	"testing"

	"github.com/cdvelop/input"
)

var (
	modelText = input.Text()

	dataText = map[string]struct {
		inputData       string
		skip_validation bool
		expected        string
	}{
		"nombre correcto con punto?":         {"Dr. Maria Jose Diaz Cadiz", false, ""},
		"no tilde ":                          {"peréz del rozal", false, "é con tilde no permitida"},
		"texto con ñ ":                       {"Ñuñez perez", false, ""},
		"texto correcto + 3 caracteres ":     {"hola", false, ""},
		"texto correcto 3 caracteres ":       {"los", false, ""},
		"oración ok ":                        {"hola que tal", false, ""},
		"solo Dato numérico permitido?":      {"100", false, ""},
		"con caracteres y coma ":             {"los,true, vengadores", false, ""},
		"sin data ok":                        {"", false, "tamaño mínimo 2 caracteres"},
		"un carácter numérico ":              {"8", false, "tamaño mínimo 2 caracteres"},
		"palabra mas numero permitido ":      {"son 4 bidones", false, ""},
		"con paréntesis y numero ":           {"son 4 (4 bidones)", false, ""},
		"con solo paréntesis ":               {"son (bidones)", false, ""},
		"palabras y numero":                  {"apellido Actualizado 1", false, ""},
		"palabra con slash?":                 {" estos son \\n los podria", false, "carácter \\ no permitido"},
		"nombre de archivos separados por ,": {"dino.png, gatito.jpeg", false, ""},
	}
)

func Test_TagText(t *testing.T) {
	tag := modelText.BuildContainerView("1", "name", true)
	if tag == "" {
		log.Fatalln("ERROR NO TAG RENDERING ")
	}
}

func Test_InputText(t *testing.T) {
	for prueba, data := range dataText {
		t.Run((prueba + data.inputData), func(t *testing.T) {
			err := modelText.ValidateField(data.inputData, data.skip_validation)

			var err_str string
			if err != nil {
				err_str = err.Error()
			}

			if err_str != data.expected {
				log.Println(prueba)
				log.Fatalf("resultado: [%v] expectativa: [%v]\n%v", err, data.expected, data.inputData)
			}
		})
	}
}

func Test_GoodInputText(t *testing.T) {
	for _, data := range modelText.GoodTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelText.ValidateField(data, false); ok != nil {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_GoodInputTextFirsNames(t *testing.T) {
	for _, data := range modelText.GoodTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelText.ValidateField(data, false); ok != nil {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_WrongInputText(t *testing.T) {
	for _, data := range modelText.WrongTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelText.ValidateField(data, false); ok == nil {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}
