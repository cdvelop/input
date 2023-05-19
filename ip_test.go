package input_test

import (
	"log"
	"testing"

	"github.com/cdvelop/input"
)

var (
	modelIp = input.Ip()

	dataIp = map[string]struct {
		inputData       string
		skip_validation bool
		expected        bool
	}{
		"ip correcta ":   {"192.168.1.1", false, true},
		"ip incorrecta ": {"192.168.1.1.8", false, false},
		"correcto?":      {"0.0.0.0", false, false},
		"sin data ":      {"", false, false},
	}
)

func Test_InputIp(t *testing.T) {
	for prueba, data := range dataIp {
		t.Run((prueba + data.inputData), func(t *testing.T) {
			if ok := modelIp.Validate.ValidateField(data.inputData, data.skip_validation); ok != data.expected {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_GoodInputIp(t *testing.T) {
	for _, data := range modelIp.TestData.GoodTestData("", "", true) {
		t.Run((data), func(t *testing.T) {
			if ok := modelIp.Validate.ValidateField(data, false); !ok {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_WrongInputIp(t *testing.T) {
	for _, data := range modelIp.TestData.WrongTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelIp.Validate.ValidateField(data, false); ok {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}
