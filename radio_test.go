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

func Test_TagRadio(t *testing.T) {
	tag := modelRadio.Tag.HtmlTag("1", "name", true)
	if tag == "" {
		log.Fatalln("ERROR NO TAG RENDERING ")
	}
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

func Test_RadioGender(t *testing.T) {

	modelGenderRadio := input.Radio(nil, "gender")

	genderData := map[string]struct {
		inputData       string
		skip_validation bool
		expected        bool
	}{
		"f Dato en minúscula correcto": {"f", false, true},
		"m Dato en minúscula correcto": {"m", false, true},
		"F Dato mayúscula incorrecto":  {"F", false, false},
		"M Dato mayúscula incorrecto":  {"M", false, false},
		"Dato existe?":                 {"1", false, false},
		"data ok?":                     {"0", false, false},
		"numero ok?":                   {"20", false, false},
		"data correcta?":               {"", false, false},
	}

	for prueba, data := range genderData {
		t.Run((prueba), func(t *testing.T) {
			if ok := modelGenderRadio.Validate.ValidateField(data.inputData, data.skip_validation); ok != data.expected {
				log.Fatalf("se obtuvo [%v] y se esperaba [%v]\n[%v]", ok, data.expected, data)
			}
		})
	}
}

func Test_GoodInputRadio(t *testing.T) {
	for _, data := range modelRadio.TestData.GoodTestData() {
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
