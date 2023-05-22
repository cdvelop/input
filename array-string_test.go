package input

import (
	"fmt"
	"testing"
)

var (
	dataCombine = map[string]struct {
		random bool
		arrays [][]string
	}{
		"no data": {arrays: [][]string{}},

		"solo un array": {
			arrays: [][]string{
				{"el", "la", "los"},
			}},

		"2 array, impar salida aleatoria": {
			random: true,
			arrays: [][]string{
				{"el", "la", "los"},
				{"gato", "perro"},
			}},
		"5 array, variado": {
			arrays: [][]string{
				{"el", "la", "los"},
				{"gato", "perro", "oso", "lagarto"}, //key 2
				{"es", "esta", "se"},                //key 3
				{"puso", "ido", "malo"},
				{"ayer", "hoy", "nose", "tal vez"}, //key 1
			}},
		"3 array, par salida desordenada": {
			random: true,
			arrays: [][]string{
				{"el", "la"},
				{"gato", "perro"},
				{"es", "esta"},
			}},
		"2 array, par": {
			arrays: [][]string{
				{"el", "la"},
				{"chancho", "vaca"},
			}},
	}
)

func Test_CombineStringArray(t *testing.T) {

	for prueba, data := range dataCombine {
		t.Run((prueba), func(t *testing.T) {
			var total_combinations int

			if len(data.arrays) != 0 {
				total_combinations = 1
			}

			array_respond := combineStringArray(data.random, data.arrays...)

			for _, array_in := range data.arrays {
				total_combinations *= len(array_in)
			}

			var count = make(map[string]struct{}, total_combinations)

			for _, v := range array_respond {
				count[v] = struct{}{}
			}

			if len(count) != total_combinations {

				for n, v := range array_respond {
					fmt.Printf("%v: [%v]\n", n, v)
				}
				// fmt.Println("COMBINACIONES", total_combinations, "SALIDA: ", len(count))
				t.Fatal()
			}
		})
	}
}
