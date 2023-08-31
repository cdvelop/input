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
		expected        string
	}{
		"D Dato correcto":                {"D", false, ""},
		"V Dato correcto":                {"V", false, ""},
		"d Dato en minúscula incorrecto": {"d", false, "valor d no corresponde a botón radio"},
		"v Dato en minúscula incorrecto": {"v", false, "valor v no corresponde a botón radio"},
		"Dato existe?":                   {"1", false, ""},
		"data ok?":                       {"0", false, "valor 0 no corresponde a botón radio"},
		"numero ok?":                     {"20", false, ""},
		"data correcta?":                 {"", false, ""},
	}
)

type radio struct{}

func (radio) SourceData() map[string]string {
	return map[string]string{"": "sin data", "1": "1 min.", "D": "Dama", "V": "Varón", "20": "20 min"}
}

func Test_TagRadio(t *testing.T) {
	tag := modelRadio.Tag.HtmlTag("1", "name", true)
	if tag == "" {
		log.Fatalln("ERROR NO TAG RENDERING ")
	}
}

func Test_RadioButton(t *testing.T) {
	for prueba, data := range TestData {
		t.Run((prueba), func(t *testing.T) {
			err := modelRadio.Validate.ValidateField(data.inputData, data.skip_validation)
			var resp string
			if err != nil {
				resp = err.Error()
			}

			if resp != data.expected {
				log.Println(prueba)
				log.Fatalf("resultado: [%v] expectativa: [%v]\n%v", resp, data.expected, data.inputData)
			}
		})
	}
}

func Test_RadioGender(t *testing.T) {

	modelGenderRadio := input.Radio(nil, "gender")

	genderData := map[string]struct {
		inputData       string
		skip_validation bool
		expected        string
	}{
		"f Dato en minúscula correcto": {"f", false, ""},
		"m Dato en minúscula correcto": {"m", false, ""},
		"F Dato mayúscula incorrecto":  {"F", false, "valor F no corresponde a botón radio"},
		"M Dato mayúscula incorrecto":  {"M", false, "valor M no corresponde a botón radio"},
		"Dato existe?":                 {"1", false, "valor 1 no corresponde a botón radio"},
		"data ok?":                     {"0", false, "valor 0 no corresponde a botón radio"},
		"numero ok?":                   {"20", false, "valor 20 no corresponde a botón radio"},
		"data correcta?":               {"", false, "selección requerida"},
	}

	for prueba, data := range genderData {
		t.Run((prueba), func(t *testing.T) {
			err := modelGenderRadio.Validate.ValidateField(data.inputData, data.skip_validation)
			var resp string
			if err != nil {
				resp = err.Error()
			}

			if resp != data.expected {
				log.Println(prueba)
				log.Fatalf("resultado: [%v] expectativa: [%v]\n%v", resp, data.expected, data.inputData)
			}
		})
	}
}

func Test_GoodInputRadio(t *testing.T) {
	for _, data := range modelRadio.TestData.GoodTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelRadio.Validate.ValidateField(data, false); ok != nil {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_WrongInputRadio(t *testing.T) {
	for _, data := range modelRadio.TestData.WrongTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelRadio.Validate.ValidateField(data, false); ok == nil {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}
