package input_test

import (
	"log"
	"testing"

	"github.com/cdvelop/input"
)

var (
	modelHour = input.Hour()

	dataHour = map[string]struct {
		inputData       string
		skip_validation bool
		expected        bool
	}{
		"correcto":    {"23:59", false, true},
		"correcto 00": {"00:00", false, true},
		"correcto 12": {"12:00", false, true},

		"incorrecto 24":       {"24:00", false, false},
		"incorrecto sin data": {"", false, false},
		"incorrecto carácter": {"13-34", false, false},
	}
)

func Test_InputHour(t *testing.T) {
	for prueba, data := range dataHour {
		t.Run((prueba + " " + data.inputData), func(t *testing.T) {
			if ok := modelHour.Validate.ValidateField(data.inputData, data.skip_validation); ok != data.expected {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_GoodInputHour(t *testing.T) {
	for _, data := range modelHour.TestData.GoodTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelHour.Validate.ValidateField(data, false); !ok {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_WrongInputHour(t *testing.T) {
	for _, data := range modelHour.TestData.WrongTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelHour.Validate.ValidateField(data, false); ok {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}
