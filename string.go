package input

type strings struct{}

func String() *strings {
	return &strings{}
}

// Contains verifica si la cadena 'search' está presente en 'text'
func (strings) Contains(text, search string) int {
	// Si la cadena de búsqueda es una cadena vacía, no puede haber coincidencias
	if search == "" {
		return 0
	}

	// Obtén la longitud de la cadena de búsqueda
	searchLen := len(search)

	// Inicializa el contador de coincidencias
	count := 0

	// Recorre el texto y cuenta la cantidad de coincidencias
	for i := 0; i <= len(text)-searchLen; i++ {
		if text[i:i+searchLen] == search {
			count++
		}
	}

	// Devuelve la cantidad de coincidencias encontradas
	return count
}

func (strings) Replace(text, old, newStr string) (result string) {

	for i := 0; i < len(text); i++ {
		// Buscar la ocurrencia de la palabra antigua en el texto
		if i+len(old) <= len(text) && text[i:i+len(old)] == old {
			// Agregar la nueva palabra al resultado
			result += newStr
			// Saltar la longitud de la palabra antigua en el texto original
			i += len(old) - 1
		} else {
			// Agregar el carácter actual al resultado
			result += string(text[i])
		}
	}

	return result
}

func (strings) TrimSuffix(text, suffix string) string {
	if len(text) < len(suffix) || text[len(text)-len(suffix):] != suffix {
		return text
	}
	return text[:len(text)-len(suffix)]
}

// ej: "texto1,texto2" "," = []string{texto1,texto2}
func (strings) Split(data, separator string) (result []string) {

	if len(data) >= 3 {

		start := 0

		for i := 0; i < len(data); i++ {
			if data[i:i+len(separator)] == separator {
				result = append(result, data[start:i])
				start = i + len(separator)
				i += len(separator) - 1
			}
		}

		result = append(result, data[start:])
	} else {
		return []string{data}
	}

	return
}

// Eliminar espacios en blanco al principio y al final
func (strings) Trim(text string) string {
	// Eliminar espacios al principio
	start := 0
	for start < len(text) && text[start] == ' ' {
		start++
	}

	// Eliminar espacios al final y al final de cada línea
	end := len(text) - 1
	for end >= 0 && (text[end] == ' ' || text[end] == '\n' || text[end] == '\t') {
		end--
	}

	// Caso especial: cadena vacía
	if start > end {
		return ""
	}

	// Devolver la subcadena sin espacios
	return text[start : end+1]
}

// solo a minúscula texto del alfabeto con ñ
func (s strings) ToLowerCase(new string) string {

	var out string

	for _, c := range new {

		if l, exist := s.LettersUpperLowerCase(true)[c]; exist {
			out += string(l)
		} else {
			out += string(c)
		}
	}

	return out
}

var letters = map[rune]rune{
	'A': 'a', 'B': 'b', 'C': 'c', 'D': 'd', 'E': 'e', 'F': 'f', 'G': 'g', 'H': 'h', 'I': 'i',
	'J': 'j', 'K': 'k', 'L': 'l', 'M': 'm', 'N': 'n', 'O': 'o', 'P': 'p', 'Q': 'q', 'R': 'r',
	'S': 's', 'T': 't', 'U': 'u', 'V': 'v', 'W': 'w', 'X': 'x', 'Y': 'y', 'Z': 'z',
}

// ej: A:a, B:b.. with_ñ = Ñ:ñ
func (strings) LettersUpperLowerCase(with_ñ ...bool) map[rune]rune {

	for _, ñ := range with_ñ {
		if ñ {
			letters['Ñ'] = 'ñ'
			return letters
		}
	}

	// Si no se incluye la letra "Ñ", se elimina del mapa
	delete(letters, 'Ñ')

	return letters
}
