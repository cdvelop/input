package input

import (
	"strconv"
)

type Permitted struct {
	Letters    bool
	Tilde      bool
	Numbers    bool
	Characters []rune //ej: '\','/','@'
	Minimum    int    //caracteres min ej 2 "lo" ok default 0 no defined
	Maximum    int    //caracteres max ej 1 "l" ok default 0 no defined
}

func (p Permitted) Validate(text string) (err string) {

	if p.Minimum != 0 {
		if len(text) < p.Minimum {
			return "tamaño mínimo " + strconv.Itoa(p.Minimum) + " caracteres"
		}
	}

	if p.Maximum != 0 {
		if len(text) > p.Maximum {
			return "tamaño máximo " + strconv.Itoa(p.Maximum) + " caracteres"
		}
	}

	for _, char := range text {

		if p.Letters {
			// fmt.Printf("Letters [%c]\n", char)
			if !valid_letters[char] {
				err = string(char) + " no es una letra"
			} else {
				err = ""
				continue
			}
		}

		if p.Tilde {
			if !valid_tilde[char] {
				err = "tilde " + string(char) + " no soportada"
			} else {
				err = ""
				continue
			}
		}

		if p.Numbers {
			// fmt.Printf("Number [%c]\n", char)
			if !valid_number[char] {
				if char == ' ' {
					err = "espacios en blanco no permitidos"
				} else {
					err = string(char) + " no es un numero"
				}

			} else {
				err = ""
				continue
			}
		}

		if len(p.Characters) != 0 {
			var found bool
			for _, c := range p.Characters {
				if c == char {
					found = true
					break
				}
			}

			if found {
				// fmt.Printf("Character ok: [%c]\n", char)
				err = ""
				continue
			} else {

				if char == ' ' {
					return "espacios en blanco no permitidos"
				} else if valid_tilde[char] {
					return string(char) + " con tilde no permitida"
				}

				return "carácter " + string(char) + " no permitido"
			}
		}

		if err != "" {
			return
		}
	}

	return err
}

// Define un mapa de caracteres válidos
var valid_letters = map[rune]bool{
	'a': true, 'b': true, 'c': true, 'd': true, 'e': true, 'f': true, 'g': true, 'h': true, 'i': true,
	'j': true, 'k': true, 'l': true, 'm': true, 'n': true, 'o': true, 'p': true, 'q': true, 'r': true,
	's': true, 't': true, 'u': true, 'v': true, 'w': true, 'x': true, 'y': true, 'z': true,
	'ñ': true,

	'A': true, 'B': true, 'C': true, 'D': true, 'E': true, 'F': true, 'G': true, 'H': true, 'I': true,
	'J': true, 'K': true, 'L': true, 'M': true, 'N': true, 'O': true, 'P': true, 'Q': true, 'R': true,
	'S': true, 'T': true, 'U': true, 'V': true, 'W': true, 'X': true, 'Y': true, 'Z': true,
	'Ñ': true,
}

var valid_tilde = map[rune]bool{
	'á': true, 'é': true, 'í': true, 'ó': true, 'ú': true,
}

var valid_number = map[rune]bool{
	'0': true, '1': true, '2': true, '3': true, '4': true, '5': true, '6': true, '7': true, '8': true, '9': true,
}

func (p Permitted) ValidateField(data_in string, skip_validation bool, options ...string) (err string) {
	if !skip_validation {
		return p.Validate(data_in)
	}
	return ""
}
