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
	modelSelect = input.SelecTag(data{})

	dataSelect = map[string]struct {
		inputData       string
		skip_validation bool
		expected        bool
	}{
		"una credencial ok?":  {"1", false, true},
		"otro numero ok?":     {"3", false, true},
		"0 existe?":           {"0", false, false},
		"-1 valido?":          {"-1", false, false},
		"carácter permitido?": {"%", false, false},
		"con data?":           {"", false, false},
		"sin espacios?":       {"luis ", false, false},
	}
)

func Test_Select(t *testing.T) {
	for prueba, data := range dataSelect {
		t.Run((prueba + " " + data.inputData), func(t *testing.T) {
			if ok := modelSelect.Validate.ValidateField(data.inputData, data.skip_validation); ok != data.expected {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}
func Test_GoodInputSelect(t *testing.T) {
	for _, data := range modelSelect.TestData.GoodTestData("", "", true) {
		t.Run((data), func(t *testing.T) {
			if ok := modelSelect.Validate.ValidateField(data, false); !ok {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_WrongInputSelect(t *testing.T) {
	for _, data := range modelSelect.TestData.WrongTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelSelect.Validate.ValidateField(data, false); ok {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}
