package input

import "fmt"

func (d dateAge) HtmlTAG(id, field_name string, allow_skip_completed bool) string {

	var valide string
	if !allow_skip_completed { //si cadena de validación no esta vacía
		valide = ` pattern="` + d.Pattern + `" required`
	}

	tag := `<label class="age-number"><input data-name="age-number" type="number" min="0" max="150" title="Campo Informativo"></label>`

	tag += fmt.Sprintf(`<label class="age-date"><input type="date" data-name="%v" name="%v" title="%v" %v></label>`,
		field_name, field_name, d.Title, valide)

	return tag
}
