package input_test

import (
	"log"
	"testing"

	"github.com/cdvelop/input"
)

var (
	dataIp = map[string]struct {
		inputData       string
		skip_validation bool
		expected        string
	}{
		"IPv4 ok":        {"192.168.1.1", false, ""},
		"IPv6 ok":        {"2001:0db8:85a3:0000:0000:8a2e:0370:7334", false, ""},
		"ip incorrecta ": {"192.168.1.1.8", false, "formato IPv4 no valida"},
		"correcto?":      {"0.0.0.0", false, "ip de ejemplo no valida"},
		"sin data ":      {"", false, "version IPv4 o 6 no encontrada"},
	}
)

func Test_InputIp(t *testing.T) {
	for prueba, data := range dataIp {
		t.Run((prueba + data.inputData), func(t *testing.T) {
			err := input.Ip().Validate.ValidateField(data.inputData, data.skip_validation)

			if err != data.expected {
				log.Println(prueba)
				log.Fatalf("resultado: [%v] expectativa: [%v]\n%v", err, data.expected, data.inputData)
			}
		})
	}
}

func Test_TagIp(t *testing.T) {
	tag := input.Ip().Tag.BuildContainerView("1", "name", true)
	if tag == "" {
		log.Fatalln("ERROR NO TAG RENDERING ")
	}
}

func Test_GoodInputIp(t *testing.T) {
	for _, data := range input.Ip().TestData.GoodTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := input.Ip().Validate.ValidateField(data, false); ok != "" {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_WrongInputIp(t *testing.T) {
	for _, data := range input.Ip().TestData.WrongTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := input.Ip().Validate.ValidateField(data, false); ok == "" {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}
