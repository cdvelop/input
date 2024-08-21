package input

import (
	"reflect"
)

// Action:radio,text,number etc TagInput <input type="input"
func (a attributes) BuildHtmlTag(htmlName, customName, id, fieldName string) string {

	var open string
	var close string

	switch htmlName {
	case "textarea":
		open = `<textarea `
		close = `></textarea>`

	default:
		open = `<input type="` + htmlName + `" `
		close = `>`

	}

	result := open + `id="` + id + `" name="` + fieldName + `" data-name="` + customName + `"`

	elem := reflect.ValueOf(a)
	elemType := elem.Type()

	for i := 0; i < elem.NumField(); i++ {
		field := elem.Field(i)

		attributeName := elemType.Field(i).Name

		// Skip if the field is not of type string
		if field.Kind() != reflect.String {
			continue
		}

		fieldValue := field.Interface().(string)

		if fieldValue != "" {

			switch attributeName {

			case "Pattern":
				// if !AllowSkipCompleted {
				// 	result += ` pattern="` + a.Pattern + `"`
				// }

			case "Min":
				result += ` min="` + a.Min + `"`
			case "Max":
				result += ` max="` + a.Max + `"`

			default:
				result += ` ` + fieldValue

			}

		}

	}

	if a.Onchange == "" && a.Onkeyup == "" && a.Oninput == "" && htmlName != "hidden" {
		result += ` oninput="` + DefaultValidateFunction + `"`
	}

	if !a.AllowSkipCompleted {
		result += ` required`
	}

	result += close

	return result
}
