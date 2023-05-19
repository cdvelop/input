package input_test

import (
	"log"
	"testing"

	"github.com/cdvelop/input"
)

var (
	modelPrimaryKey = input.Pk()

	dataPrimaryKey = map[string]struct {
		inputData       string
		skip_validation bool
		expected        bool
	}{
		"id 1 correcto?":        {"1624397134562544800", false, true},
		"id 2 ok?":              {"1624397172303448900", false, true},
		"id 3 ok?":              {"1634394443466878800", false, true},
		"numero 5 correcto?":    {"5", false, true},
		"numero 45 correcto?":   {"45", false, true},
		"id con letra valido?":  {"E624397172303448900", false, false},
		"id con data completa?": {"", false, false},
		"id cero?":              {"0", false, true},
	}
)

func Test_InputPrimaryKey(t *testing.T) {

	for prueba, data := range dataPrimaryKey {
		t.Run((prueba + ": " + data.inputData), func(t *testing.T) {
			if ok := modelPrimaryKey.Validate.ValidateField(data.inputData, data.skip_validation); ok != data.expected {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_GoodInputPrimaryKey(t *testing.T) {
	for _, data := range modelPrimaryKey.TestData.GoodTestData("", "", true) {
		t.Run((data), func(t *testing.T) {
			if ok := modelPrimaryKey.Validate.ValidateField(data, false); !ok {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_WrongInputPrimaryKey(t *testing.T) {
	for _, data := range modelPrimaryKey.TestData.WrongTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelPrimaryKey.Validate.ValidateField(data, false); ok {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}
