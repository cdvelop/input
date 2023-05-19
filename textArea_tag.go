package input

import (
	"fmt"
)

func (t textArea) HtmlTAG(id, field_name string, allow_skip_completed bool) string {

	valide := `pattern="` + t.pattern + `"`
	title := `title="` + t.info + `"`

	var required string
	if !allow_skip_completed { //si cadena de validación no esta vacía
		required = ` required`
	}

	return fmt.Sprintf(`<textarea name="%v" data-name="%v" rows="%v" cols="%v" %v %v%v></textarea>`, field_name, field_name, t.Rows, t.Cols, title,
		valide, required)
}
