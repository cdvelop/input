package input_test

import (
	"log"
	"testing"

	"github.com/cdvelop/input"
)

type data struct{}

func (data) SourceData() map[string]string {
	return map[string]string{"1": "Admin", "2": "editor", "3": "visitante"}
}

var (
	modelSelect = input.SelecTag("UserType", data{})

	dataSelect = map[string]struct {
		inputData       string
		skip_validation bool
		expected        string
	}{
		"una credencial ok?":  {"1", false, ""},
		"otro numero ok?":     {"3", false, ""},
		"0 existe?":           {"0", false, "valor 0 no corresponde al select"},
		"-1 valido?":          {"-1", false, "valor -1 no corresponde al select"},
		"carácter permitido?": {"%", false, "valor % no corresponde al select"},
		"con data?":           {"", false, "selección requerida"},
		"sin espacios?":       {"luis ", false, "valor luis  no corresponde al select"},
	}
)

func Test_TagSelect(t *testing.T) {
	tag := modelSelect.Tag.BuildContainerView("1", "name", true)
	if tag == "" {
		log.Fatalln("ERROR NO TAG RENDERING ")
	}
}

func Test_Select(t *testing.T) {
	for prueba, data := range dataSelect {
		t.Run((prueba + " " + data.inputData), func(t *testing.T) {
			err := modelSelect.Validate.ValidateField(data.inputData, data.skip_validation)

			if err != data.expected {
				log.Println(prueba)
				log.Fatalf("resultado: [%v] expectativa: [%v]\n%v", err, data.expected, data.inputData)
			}
		})
	}
}
func Test_GoodInputSelect(t *testing.T) {
	for _, data := range modelSelect.TestData.GoodTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelSelect.Validate.ValidateField(data, false); ok != "" {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_WrongInputSelect(t *testing.T) {
	for _, data := range modelSelect.TestData.WrongTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelSelect.Validate.ValidateField(data, false); ok == "" {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}
