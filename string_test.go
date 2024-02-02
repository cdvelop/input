package input_test

import (
	"testing"

	"github.com/cdvelop/input"
)

func TestContains(t *testing.T) {

	var testCases = map[string]struct {
		text     string
		search   string
		expected int
	}{
		"Caso1": {
			text:     "Hola, mundo!",
			search:   "mundo",
			expected: 1,
		},
		"Caso2": {
			text:     "Hola, mundo!",
			search:   "golang",
			expected: 0,
		},
		"Caso3": {
			text:     "Hola, mundo!",
			search:   "",
			expected: 0,
		},
		"Caso4": {
			text:     "Hola",
			search:   "Hola, mundo!",
			expected: 0,
		},
		"Caso5": {
			text:     "abracadabra",
			search:   "abra",
			expected: 2,
		},
		"Caso6": {
			text:     "abracadabra",
			search:   "bra",
			expected: 2,
		},
		"Caso7": {
			text:     "abra,cadabra",
			search:   ",",
			expected: 1,
		},
		"Caso8": {
			text:     "(abraLol,*?¡¡",
			search:   "Lol",
			expected: 1,
		},
		"Caso9 ": {
			text:     "(abraLol,*?¡¡",
			search:   "LoL",
			expected: 0,
		},
		"Caso10 ": {
			text:     "(¡ab¡raLol,*?¡¡",
			search:   "¡",
			expected: 4,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {

			expected := input.String().Contains(tc.text, tc.search)
			if expected != tc.expected {
				t.Errorf("Error: Se esperaba %v, pero se obtuvo %v. Texto: %s, Búsqueda: %s", tc.expected, expected, tc.text, tc.search)
			}
		})
	}
}

func TestReplace(t *testing.T) {
	tests := []struct {
		input, old, newStr, expected string
	}{
		{"Este es un ejemplo de texto de prueba.", "ejemplo", "cambio", "Este es un cambio de texto de prueba."},
		{"Hola mundo!", "mundo", "Gophers", "Hola Gophers!"},
		{"abc abc abc", "abc", "123", "123 123 123"},
		{"abc", "xyz", "123", "abc"},
		{"", "", "123", ""},
		{"abcdabcdabcd", "cd", "12", "ab12ab12ab12"},
		{"palabra, punto,", ",", ".", "palabra. punto."},
	}

	for _, test := range tests {
		result := input.String().Replace(test.input, test.old, test.newStr)
		if result != test.expected {
			t.Errorf("Para input '%s', old '%s', new '%s', esperado '%s', pero obtenido '%s'", test.input, test.old, test.newStr, test.expected, result)
		}
	}
}

func TestTrimSuffix(t *testing.T) {
	tests := []struct {
		input, suffix, expected string
	}{
		{"hello.txt", ".txt", "hello"},
		{"example", "123", "example"},
		{"file.txt.txt", ".txt", "file.txt"},
		{"", "", ""},
		{"abc", "xyz", "abc"},
	}

	for _, test := range tests {
		result := input.String().TrimSuffix(test.input, test.suffix)
		if result != test.expected {
			t.Errorf("Para input '%s', suffix '%s', esperado '%s', pero obtenido '%s'", test.input, test.suffix, test.expected, result)
		}
	}
}

func TestStringSplit(t *testing.T) {
	testCases := []struct {
		data      string
		separator string
		expected  []string
	}{
		{"texto1,texto2", ",", []string{"texto1", "texto2"}},
		{"apple,banana,cherry", ",", []string{"apple", "banana", "cherry"}},
		{"one.two.three.four", ".", []string{"one", "two", "three", "four"}},
		{"hello world", " ", []string{"hello", "world"}},
		{"hello. world", ".", []string{"hello", " world"}},
		{"h.", ".", []string{"h."}},
	}

	for _, tc := range testCases {
		result := input.String().Split(tc.data, tc.separator)

		if !areStringSlicesEqual(result, tc.expected) {
			t.Errorf("StringSplit(%s, %s) = %v; expected %v", tc.data, tc.separator, result, tc.expected)
		}
	}
}

func areStringSlicesEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestTrim(t *testing.T) {
	tests := []struct {
		input, expected string
	}{
		{"  hello world  ", "hello world"},
		{"abc123", "abc123"},
		{"  trim me  ", "trim me"},
		{"", ""},
		{"  ", ""},

		{"    mucho espacio\n\n\t\tcon salto\n\n\t\tde linea     \n\t\t\t\t\t\t\n\t\t\t\t", "mucho espacio\n\n\t\tcon salto\n\n\t\tde linea"},

		{`    mucho espacio
		
		con salto

		de linea     
		              
		
		`, `mucho espacio
		
		con salto

		de linea`},
	}

	for _, test := range tests {
		result := input.String().Trim(test.input)
		if result != test.expected {
			t.Errorf("Para input '%s', esperado '%s', pero obtenido '%s'", test.input, test.expected, result)
		}
	}
}
