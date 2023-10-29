package input_test

import (
	"log"
	"testing"

	"github.com/cdvelop/input"
)

type checkData struct{}

func (c checkData) SourceData() map[string]string {
	return map[string]string{"1": "Admin", "2": "editor", "3": "visitante"}
}

var (
	modelCheck = input.CheckBox("TypeUser", checkData{})

	datacheck = map[string]struct {
		inputData       string
		skip_validation bool
		expected        string
	}{
		"una credencial ok?":         {modelCheck.TestData.GoodTestData()[0], false, ""},
		"editor y admin ok?":         {"1,2", false, ""},
		"todas las credenciales ok?": {`1,3`, false, ""},
		"0 existe?":                  {"0", false, "valor 0 no corresponde al checkbox"},
		"-1 valido?":                 {"-1", false, "valor -1 no corresponde al checkbox"},
		"todas existentes?":          {"1,5", false, "valor 5 no corresponde al checkbox"},
		"con data?":                  {"", false, "selección requerida"},
		"sin espacios?":              {"luis ,true, 3", false, "valor luis  no corresponde al checkbox"},
	}
)

func Test_check(t *testing.T) {
	for prueba, data := range datacheck {
		t.Run((prueba + " " + data.inputData), func(t *testing.T) {
			err := modelCheck.Validate.ValidateField(data.inputData, data.skip_validation)
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

func Test_TagCheck(t *testing.T) {
	tag := modelCheck.Tag.HtmlTag("1", "name", true)
	if tag == "" {
		log.Fatalln("ERROR NO TAG RENDERING ")
	}
}

func Test_GoodInputCheck(t *testing.T) {
	for _, data := range modelCheck.TestData.GoodTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelCheck.Validate.ValidateField(data, false); ok != nil {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_WrongInputCheck(t *testing.T) {
	for _, data := range modelCheck.TestData.WrongTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelCheck.Validate.ValidateField(data, false); ok == nil {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}
