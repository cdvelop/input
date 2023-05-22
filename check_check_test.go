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
	// newData = checkData{}

	m_check = input.Check(checkData{}, false)

	datacheck = map[string]struct {
		inputData       string
		skip_validation bool
		result          bool
	}{
		"una credencial ok?":         {m_check.TestData.GoodTestData()[0], false, true},
		"editor y admin ok?":         {"1,2", false, true},
		"todas las credenciales ok?": {`1,3`, false, true},
		"0 existe?":                  {"0", false, false},
		"-1 valido?":                 {"-1", false, false},
		"todas existentes?":          {"1,5", false, false},
		"con data?":                  {"", false, false},
		"sin espacios?":              {"luis ,true, 3", false, false},
	}
)

func Test_check(t *testing.T) {
	for prueba, data := range datacheck {
		t.Run((prueba + " " + data.inputData), func(t *testing.T) {
			if ok := m_check.Validate.ValidateField(data.inputData, data.skip_validation); ok != data.result {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_GoodInputCheck(t *testing.T) {
	for _, data := range m_check.TestData.GoodTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := m_check.Validate.ValidateField(data, false); !ok {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_WrongInputCheck(t *testing.T) {
	for _, data := range m_check.TestData.WrongTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := m_check.Validate.ValidateField(data, false); ok {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}
