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

	Autocomplete string

	Rows string //filas ej 4,5,6
	Cols string //columnas ej 50,80

	Step    string
	Oninput string // ej: "miFuncion()" = oninput="miFuncion()"

	Value string
}
