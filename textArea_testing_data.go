package input

// option prescription,
func (t textArea) GoodTestData() (out []string) {
	phrase := []string{"hola: esto, es. la - prueba 10",
		"soy ñato o Ñato (aqui)", "son dos examenes", "costo total es de $23.230. pesos",
	}

	placeholder := String().ToLowerCase(t.PlaceHolder)

	switch {
	case String().Contains(placeholder, "nombre y apellido") != 0:
		return phrase

	case String().Contains(placeholder, "diagnostic") != 0:
		return permutation(discomforts, prepositions, body_parts)

	case String().Contains(placeholder, "prescription") != 0:
		return permutation(prescription, prepositions, body_parts)

	default:
		return permutation(phrase, prepositions, hours)
	}
}

func (t textArea) WrongTestData() (out []string) {

	out = []string{
		"] son",
		" ", "& ", "SELECT * FROM", "=",
	}

	return
}
