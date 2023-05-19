package input

import (
	"fmt"
)

func (d dayWord) HtmlTAG(id, field_name string, allow_skip_completed bool) string {
	var valide string
	if !allow_skip_completed { //si cadena de validación no esta vacía
		valide = ` pattern="` + patternMonthDay + `" required`
	}

	tag := fmt.Sprintf(`<label class="date-spanish"><input`+id+` data-spanish=""  type="%v" name="%v" title="%v" %v></label>`,
		d.HtmlName(), field_name, d.Title, valide)
	return tag
}
