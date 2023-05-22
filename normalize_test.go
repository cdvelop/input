package input

import (
	"testing"
)

var (
	dataAcent = map[string]struct {
		input  string
		result string
	}{
		"3 tildes": {"mi, Nómbré í", "mi, Nombre i"},
		"1 tilde":  {"César", "Cesar"},
	}
)

func Test_RemoveAcent(t *testing.T) {
	for prueba, data := range dataAcent {
		t.Run((prueba), func(t *testing.T) {
			if res := RemoveAcent(data.input); res != data.result {
				t.Fatalf("err input [%v] resultado [%v]", data.input, res)
			}
		})
	}
}

var (
	dataTextNormalize = map[string]struct {
		input  string
		result string
	}{
		"3 tildes":               {"mi, Nómbré í", "mi, Nombre i"},
		"1 tilde":                {"César", "Cesar"},
		"cambio vació":           {"", ""},
		"un solo carácter":       {"Í", "I"},
		"espacio al final":       {"Junán ", "Junan"},
		"espacio inicio y final": {" Junán ", "Junan"},
		"Ñ mayúscula":            {" Ñandús ", "Nandus"},
	}
)

func Test_NormalizeText(t *testing.T) {
	for prueba, data := range dataTextNormalize {
		t.Run((prueba), func(t *testing.T) {
			NormalizeTextData(&data.input)
			if data.input != data.result {
				t.Fatalf("err input [%v] se esperaba [%v]", data.input, data.result)
			}
		})
	}
}
