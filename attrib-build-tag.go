package input

import (
	"reflect"
)

// Type:radio,text,number etc TagInput <input type="input"
func (a attributes) Build(html_name, input_name, id, field_name string, allow_skip_completed bool) string {

	result := `<input type="` + html_name + `" id="` + id + `" name="` + field_name + `" data-name="` + input_name + `"`

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

	result += ">"

	return result
}
