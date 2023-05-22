package input_test

import (
	"log"
	"testing"

	"github.com/cdvelop/input"
)

var (
	modelMonthDay = input.MonthDay()

	dataMonthDay = map[string]struct {
		inputData       string
		skip_validation bool
		expected        bool
	}{
		"dia ok?":                     {"01", false, true},
		"mes ok?":                     {"31", false, true},
		"fecha recortada sin año ok?": {"31-01", false, false},
		"correcto ?":                  {"1-1", false, false},
		"incorrecto ":                 {"2002-12-03", false, false},
		"formato incorrecto ":         {"21/12", false, false},
		"data incorrecta ":            {"0000-00-00", false, false},
		"toda la data correcta?":      {"", false, false},
	}
)

func Test_InputMonthDay(t *testing.T) {
	for prueba, data := range dataMonthDay {
		t.Run((prueba + data.inputData), func(t *testing.T) {
			if ok := modelMonthDay.Validate.ValidateField(data.inputData, data.skip_validation); ok != data.expected {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_GoodInputMonthDay(t *testing.T) {
	for _, data := range modelMonthDay.TestData.GoodTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelMonthDay.Validate.ValidateField(data, false); !ok {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_WrongInputMonthDay(t *testing.T) {
	for _, data := range modelMonthDay.TestData.WrongTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelMonthDay.Validate.ValidateField(data, false); ok {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}
