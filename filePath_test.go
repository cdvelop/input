package input_test

import (
	"log"
	"testing"

	"github.com/cdvelop/input"
)

var (
	filePathTestData = map[string]struct {
		inputData       string
		skip_validation bool
		expected        bool
	}{
		"dirección correcta":                               {".\\files\\1234\\", false, true},
		"dirección correcta sin punto inicio?":             {"\\files\\1234\\", false, false},
		"ruta relativa con directorios":                    {".\\files\\1234\\", false, true},
		"tres rutas separadas por comas":                   {`.\\files\\1234\\,.\\files\\5678\\,.\\images\\ok\\`, false, false},
		"ruta relativa sin punto de inicio":                {"\\files\\1234\\", false, false},
		"ruta absoluta en Linux":                           {"/home/user/files/", false, true},
		"ruta absoluta sin punto de inicio":                {"/files/1234/", false, true},
		"ruta relativa sin directorios ok?":                {".\\", false, false},
		"ruta relativa sin barra final":                    {"./files/1234", false, true},
		"ruta relativa con barra final":                    {"./files/1234/", false, true},
		"ruta con nombre de archivo":                       {".\\files\\1234\\archivo.txt", false, true},
		"ruta con nombres de directorio con guiones bajos": {".\\mi_directorio\\sub_directorio\\", false, true},
		"un numero es una ruta valida?":                    {"5", false, true},
		"una sola palabra es una ruta valida?":             {"ruta", false, true},
	}
)

func Test_Check(t *testing.T) {

	for prueba, data := range filePathTestData {
		t.Run((prueba + " " + data.inputData), func(t *testing.T) {
			if ok := input.FilePath().Validate.ValidateField(data.inputData, data.skip_validation); ok != data.expected {
				log.Fatalf("resultado [%v] data [%v]", ok, data.inputData)
			}
		})
	}
}

func Test_GoodInputFilePath(t *testing.T) {
	for _, data := range input.FilePath().GoodTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := input.FilePath().Validate.ValidateField(data, false); !ok {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_WrongInputFilePath(t *testing.T) {
	for _, data := range input.FilePath().WrongTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := input.FilePath().Validate.ValidateField(data, false); ok {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}
