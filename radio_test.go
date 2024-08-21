package input_test

import (
	"log"
	"testing"

	"github.com/cdvelop/input"
)

var (
	modelRadio = input.Radio("TestRadio", radio{})
	// modelRadio =  input_radio.radio{VALUES: }

	TestData = map[string]struct {
		inputData string

		expected string
	}{
		"D Dato correcto":                {"D", ""},
		"V Dato correcto":                {"V", ""},
		"d Dato en minúscula incorrecto": {"d", "valor d no corresponde a botón radio"},
		"v Dato en minúscula incorrecto": {"v", "valor v no corresponde a botón radio"},
		"Dato existe?":                   {"1", ""},
		"data ok?":                       {"0", "valor 0 no corresponde a botón radio"},
		"numero ok?":                     {"20", ""},
		"data correcta?":                 {"", ""},
	}
)

type radio struct{}

func (radio) SourceData() map[string]string {
	return map[string]string{"": "sin data", "1": "1 min.", "D": "Dama", "V": "Varón", "20": "20 min"}
}

func Test_TagRadio(t *testing.T) {
	tag := modelRadio.BuildInputHtml("1", "name")
	if tag == "" {
		log.Fatalln("ERROR NO TAG RENDERING ")
	}
}

func Test_RadioButton(t *testing.T) {
	for prueba, data := range TestData {
		t.Run((prueba), func(t *testing.T) {
			err := modelRadio.ValidateInput(data.inputData)

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

func Test_RadioGender(t *testing.T) {

	modelGenderRadio := input.RadioGender()

	genderData := map[string]struct {
		inputData string

		expected string
	}{
		"f Dato en minúscula correcto": {"f", ""},
		"m Dato en minúscula correcto": {"m", ""},
		"F Dato mayúscula incorrecto":  {"F", "valor F no corresponde a botón radio"},
		"M Dato mayúscula incorrecto":  {"M", "valor M no corresponde a botón radio"},
		"Dato existe?":                 {"1", "valor 1 no corresponde a botón radio"},
		"data ok?":                     {"0", "valor 0 no corresponde a botón radio"},
		"numero ok?":                   {"20", "valor 20 no corresponde a botón radio"},
		"data correcta?":               {"", "selección requerida"},
	}

	for prueba, data := range genderData {
		t.Run((prueba), func(t *testing.T) {
			err := modelGenderRadio.ValidateInput(data.inputData)

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

func Test_GoodInputRadio(t *testing.T) {
	for _, data := range modelRadio.GoodTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelRadio.ValidateInput(data); ok != nil {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_WrongInputRadio(t *testing.T) {
	for _, data := range modelRadio.WrongTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelRadio.ValidateInput(data); ok == nil {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}
