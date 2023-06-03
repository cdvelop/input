package input_test

import (
	"log"
	"testing"

	"github.com/cdvelop/input"
)

var (
	modelMail = input.Mail()

	dataMail = map[string]struct {
		inputData       string
		skip_validation bool
		expected        bool
	}{
		"correo normal ":   {"mi.correo@mail.com", false, true},
		"correo un campo ": {"correo@mail.com", false, true},
	}
)

func Test_TagMail(t *testing.T) {
	tag := modelMail.Tag.HtmlTag("1", "name", true)
	if tag == "" {
		log.Fatalln("ERROR NO TAG RENDERING ")
	}
}

func Test_InputMail(t *testing.T) {
	for prueba, data := range dataMail {
		t.Run((prueba + data.inputData), func(t *testing.T) {
			if ok := modelMail.Validate.ValidateField(data.inputData, data.skip_validation); ok != data.expected {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_GoodInputMail(t *testing.T) {
	for _, data := range modelMail.TestData.GoodTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelMail.Validate.ValidateField(data, false); !ok {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_WrongInputMail(t *testing.T) {
	for _, data := range modelMail.TestData.WrongTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelMail.Validate.ValidateField(data, false); ok {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}
