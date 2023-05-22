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
	modelDataList = input.DataList(dataDataList{})

	dataList = map[string]struct {
		inputData       string
		skip_validation bool
		result          bool
	}{
		"una credencial ok?":  {"1", false, true},
		"otro numero ok?":     {"3", false, true},
		"0 existe?":           {"0", false, false},
		"-1 valido?":          {"-1", false, false},
		"carácter permitido?": {"%", false, false},
		"con data?":           {"", false, false},
		"sin espacios?":       {"luis ", false, false},
	}
)

func Test_DataList(t *testing.T) {
	for prueba, data := range dataList {
		t.Run((prueba + " " + data.inputData), func(t *testing.T) {
			if ok := modelSelect.Validate.ValidateField(data.inputData, data.skip_validation); ok != data.result {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}
func Test_GoodInputDataList(t *testing.T) {
	for _, data := range modelSelect.TestData.GoodTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelSelect.Validate.ValidateField(data, false); !ok {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_WrongInputDataList(t *testing.T) {
	for _, data := range modelSelect.TestData.WrongTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelSelect.Validate.ValidateField(data, false); ok {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}
