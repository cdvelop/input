package input

import (
	"reflect"
)

// Action:radio,text,number etc TagInput <input type="input"
func (a attributes) BuildHtmlTag(html_name, input_name, id, field_name string, allow_skip_completed bool) string {

	var open string
	var close string

	switch input_name {
	case "textarea":
		open = `<textarea `
		close = `></textarea>`

	default:
		open = `<input type="` + html_name + `" `
		close = `>`

	}

	result := open + `id="` + id + `" name="` + field_name + `" data-name="` + input_name + `"`

	elem := reflect.ValueOf(a)
	elemType := elem.Type()

	for i := 0; i < elem.NumField(); i++ {
		field := elem.Field(i)

		attributeName := elemType.Field(i).Name
		fieldValue := field.Interface().(string)

		if fieldValue != "" {

			switch attributeName {

			case "Pattern":
				if !allow_skip_completed {
					result += ` pattern="` + a.Pattern + `"`
				}

			case "Min":
				result += ` min="` + a.Min + `"`
			case "Max":
				result += ` max="` + a.Max + `"`

			default:
				result += ` ` + fieldValue

			}

		}

	}

	if !allow_skip_completed {
		result += ` required`
	}

	if a.Onkeyup == "" && a.Oninput == "" && html_name != "hidden" {
		result += ` onkeyup="` + DefaultValidateFunction + `"`
	}

	result += close

	return result
}
