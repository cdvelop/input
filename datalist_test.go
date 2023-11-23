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
		inputData       string
		skip_validation bool
		expected        string
	}{
		"una credencial ok?":  {"1", false, ""},
		"otro numero ok?":     {"3", false, ""},
		"0 existe?":           {"0", false, "valor 0 no permitido en datalist"},
		"-1 valido?":          {"-1", false, "valor -1 no permitido en datalist"},
		"carácter permitido?": {"%", false, "valor % no permitido en datalist"},
		"con data?":           {"", false, "selección requerida"},
		"sin espacios?":       {"luis ", false, "valor luis  no permitido en datalist"},
	}
)

func Test_DataList(t *testing.T) {
	for prueba, data := range dataList {
		t.Run((prueba + " " + data.inputData), func(t *testing.T) {
			err := modelDataList.Validate.ValidateField(data.inputData, data.skip_validation)

			if err != data.expected {
				log.Println(prueba)
				log.Fatalf("resultado: [%v] expectativa: [%v]\n%v", err, data.expected, data.inputData)
			}

		})
	}
}

func Test_TagDataList(t *testing.T) {
	tag := modelDataList.Tag.BuildContainerView("1", "name", true)
	if tag == "" {
		log.Fatalln("ERROR NO TAG RENDERING ")
	}
}

func Test_GoodInputDataList(t *testing.T) {
	for _, data := range modelDataList.TestData.GoodTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelDataList.Validate.ValidateField(data, false); ok != "" {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_WrongInputDataList(t *testing.T) {
	for _, data := range modelDataList.TestData.WrongTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelDataList.Validate.ValidateField(data, false); ok == "" {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}
