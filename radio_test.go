package input_test

import (
	"log"
	"testing"

	"github.com/cdvelop/input"
)

var (
	modelRadio = input.Radio(radio{})
	// modelRadio =  input_radio.radio{VALUES: }

	TestData = map[string]struct {
		inputData       string
		skip_validation bool
		expected        bool
	}{
		"D Dato correcto":                {"D", false, true},
		"V Dato correcto":                {"V", false, true},
		"d Dato en minúscula incorrecto": {"d", false, false},
		"v Dato en minúscula incorrecto": {"v", false, false},
		"Dato existe?":                   {"1", false, true},
		"data ok?":                       {"0", false, false},
		"numero ok?":                     {"20", false, true},
		"data correcta?":                 {"", false, true},
	}
)

type radio struct{}

func (radio) SourceData() map[string]string {
	return map[string]string{"": "sin data", "1": "1 min.", "D": "Dama", "V": "Varón", "20": "20 min"}
}

func Test_RadioButton(t *testing.T) {
	for prueba, data := range TestData {
		t.Run((prueba), func(t *testing.T) {
			if ok := modelRadio.Validate.ValidateField(data.inputData, data.skip_validation); ok != data.expected {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_GoodInputRadio(t *testing.T) {
	for _, data := range modelRadio.TestData.GoodTestData("", "", true) {
		t.Run((data), func(t *testing.T) {
			if ok := modelRadio.Validate.ValidateField(data, false); !ok {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_WrongInputRadio(t *testing.T) {
	for _, data := range modelRadio.TestData.WrongTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelRadio.Validate.ValidateField(data, false); ok {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}
