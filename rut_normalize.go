package input

import (
	"fmt"
	"strconv"
	"strings"
)

func RutNormalize(rutIn *string) bool {
	if rutIn != nil && len(*rutIn) > 3 {

		r := rut{}

		if !r.ValidateField(*rutIn, false) {
			//limpiamos elementos en blanco
			NormalizeTextData(rutIn)
			neWRutIn := strings.ToLower(*rutIn)

			if len(neWRutIn) < 2 {
				return false
			}
			neWRutIn = neWRutIn[1:] //primera prueba quitando el primer carácter
			if r.ValidateField(neWRutIn, false) {
				*rutIn = neWRutIn
				return true
			} else { //cambiemos el dígito verificador y probemos
				_, onlyRun, _ := runData(*rutIn)
				dv := DvRut(onlyRun)
				run2 := fmt.Sprintf("%v-%v", onlyRun, dv)
				// fmt.Printf("NUEVO RUN: %v\n", run2)
				if r.ValidateField(run2, false) { //ultima opción
					*rutIn = strings.ToLower(run2)
					return true
				}
			}
		} else {
			return true
		}
	}
	return false
}

func runData(runIn string) (data []string, onlyRun int, ok bool) {
	data = strings.Split(string(runIn), "-")
	var err error
	if onlyRun, _ = strconv.Atoi(data[0]); err != nil {
		return
	}
	ok = true
	return
}
