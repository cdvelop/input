package input_test

import (
	"log"
	"testing"

	"github.com/cdvelop/input"
)

var (
	modelTextOnly = input.TextOnly()

	dataTextOnly = map[string]struct {
		inputData       string
		skip_validation bool
		expected        string
	}{
		"nombre correcto con punto?":       {"Dr. Maria Jose Diaz Cadiz", false, "carácter . no permitido"},
		"palabras con tilde?":              {"María Jose Diáz Cadíz", false, "í con tilde no permitida"},
		"caracteres 47 ok?":                {"juan marcos antonio del rosario de las carmenes", false, ""},
		"tilde ok ? ":                      {"peréz del rozal", false, "é con tilde no permitida"},
		"texto con ñ?":                     {"Ñuñez perez", false, ""},
		"texto correcto + 3 caracteres ":   {"juli", false, ""},
		"texto correcto 3 caracteres ":     {"luz", false, ""},
		"oración ok ":                      {"hola que tal", false, ""},
		"Dato numérico 100 no permitido? ": {"100", false, "carácter 1 no permitido"},
		"con caracteres y coma ?":          {"los,true, vengadores", false, "carácter , no permitido"},
		"sin data ok":                      {"", false, "tamaño mínimo 3 caracteres"},
		"un carácter numérico ?":           {"8", false, "tamaño mínimo 3 caracteres"},
		"palabra mas numero permitido ?":   {"son 4 bidones", false, "carácter 4 no permitido"},
		"con paréntesis y numero ?":        {"son {4 bidones}", false, "carácter { no permitido"},
		"con solo paréntesis ?":            {"son (bidones)", false, "carácter ( no permitido"},
		"palabras y numero ?":              {"apellido Actualizado 1", false, "carácter 1 no permitido"},
		"un carácter ok?":                  {"!", false, "tamaño mínimo 3 caracteres"},
	}
)

func Test_TagTextOnly(t *testing.T) {
	tag := modelTextOnly.Tag.BuildContainerView("1", "name", true)
	if tag == "" {
		log.Fatalln("ERROR NO TAG RENDERING ")
	}
}

func Test_InputTextOnly(t *testing.T) {
	for prueba, data := range dataTextOnly {
		t.Run((prueba + data.inputData), func(t *testing.T) {
			err := modelTextOnly.Validate.ValidateField(data.inputData, data.skip_validation)

			if err != data.expected {
				log.Println(prueba)
				log.Fatalf("resultado: [%v] expectativa: [%v]\n%v", err, data.expected, data.inputData)
			}
		})
	}
}

func Test_GoodInputTextOnly(t *testing.T) {
	for _, data := range modelTextOnly.TestData.GoodTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelTextOnly.Validate.ValidateField(data, false); ok != "" {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_WrongInputTextOnly(t *testing.T) {
	for _, data := range modelTextOnly.TestData.WrongTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelTextOnly.Validate.ValidateField(data, false); ok == "" {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}
