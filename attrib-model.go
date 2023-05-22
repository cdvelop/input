package input

type attributes struct {
	DataSet string // `data-xxx="nnn"`

	Class string // clase css ej: "age"

	PlaceHolder string
	Title       string //info

	Pattern       string //caracteres para validación
	pattern_start string // para rearmado con Min y Max
	pattern_end   string // para rearmado con Max y Max

	Min string //valor mínimo
	Max string //valor máximo

	Maxlength string //ej: maxlength="12"

	Autocomplete string

	Rows string //filas ej 4,5,6
	Cols string //columnas ej 50,80

	Step    string
	Oninput string // ej: "miRealtimeFunction()" = oninput="miRealtimeFunction()"
	Onkeyup string // ej: "miNormalFuncion()" = onkeyup="miNormalFuncion()"

	// https://developer.mozilla.org/en-US/docs/Web/HTML/Attributes/accept
	// https://developer.mozilla.org/es/docs/Web/HTTP/Basics_of_HTTP/MIME_types/Common_types
	// accept="image/*"
	Accept   string
	Multiple string // multiple

	Value string
}
