package input

import (
	"log"
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func removeAcent(in string) (out string) {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFD)
	output, _, e := transform.String(t, in)
	if e != nil {
		log.Printf("\n¡¡¡Error %v al remover tilde %v\n", e, in)
	}
	return output
}

// remueve tildes,espacio blanco inicio-final y todo a minúscula
func NormalizeTextData(in *string) {
	*in = removeAcent(*in)
	*in = strings.TrimSpace(*in) //remover espacios en blanco inicial final
}

func RemoveUfeffFromString(in *string) {
	// remover caracteres al inicio del archivo texto
	*in = strings.Replace(*in, "\ufeff", "", 1)
	// fmt.Printf("%q", embed_file)
}
