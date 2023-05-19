package input

// option prescription,
func (t textArea) GoodTestData(table_name, field_name string, random bool) (out []string) {
	phrase := []string{"hola: esto, es. la - prueba 10",
		"soy ñato o Ñato (aqui)", "son dos examenes", "costo total es de $23.230. pesos",
	}

	switch field_name {
	case "phrase":
		return phrase

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
