package input_test

import (
	"log"
	"testing"

	"github.com/cdvelop/input"
)

type dataDataList struct{}

func (dataDataList) SourceData() map[string]string {
	return map[string]string{"1": "Admin", "2": "editor", "3": "visitante"}
}

var (
	modelDataList = input.DataList("UserType", dataDataList{})

	dataList = map[string]struct {
		inputData string

		expected string
	}{
		"una credencial ok?":  {"1", ""},
		"otro numero ok?":     {"3", ""},
		"0 existe?":           {"0", "valor 0 no permitido en datalist"},
		"-1 valido?":          {"-1", "valor -1 no permitido en datalist"},
		"carácter permitido?": {"%", "valor % no permitido en datalist"},
		"con data?":           {"", "selección requerida"},
		"sin espacios?":       {"luis ", "valor luis  no permitido en datalist"},
	}
)

func Test_DataList(t *testing.T) {
	for prueba, data := range dataList {
		t.Run((prueba + " " + data.inputData), func(t *testing.T) {
			err := modelDataList.ValidateInput(data.inputData)

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

func Test_TagDataList(t *testing.T) {
	tag := modelDataList.BuildInputHtml("1", "name")
	if tag == "" {
		log.Fatalln("ERROR NO TAG RENDERING ")
	}
}

func Test_GoodInputDataList(t *testing.T) {
	for _, data := range modelDataList.GoodTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelDataList.ValidateInput(data); ok != nil {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_WrongInputDataList(t *testing.T) {
	for _, data := range modelDataList.WrongTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelDataList.ValidateInput(data); ok == nil {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}
