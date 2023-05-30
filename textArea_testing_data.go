package input

import "strings"

// option prescription,
func (t textArea) GoodTestData() (out []string) {
	phrase := []string{"hola: esto, es. la - prueba 10",
		"soy ñato o Ñato (aqui)", "son dos examenes", "costo total es de $23.230. pesos",
	}

	placeholder := strings.ToLower(t.PlaceHolder)

	switch {
	case strings.Contains(placeholder, "nombre y apellido"):
		return phrase

	case strings.Contains(placeholder, "diagnostic"):
		return combineStringArray(true, discomforts, prepositions, body_parts)

	case strings.Contains(placeholder, "prescription"):
		return combineStringArray(true, prescription, prepositions, body_parts)

	default:
		return combineStringArray(true, phrase, prepositions, hours)
	}
}

func (t textArea) WrongTestData() (out []string) {

	out = []string{
		"# son",
		" ", "#", "& ", "SELECT * FROM", "=",
	}

	return
}
