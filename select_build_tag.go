package input

import (
	"fmt"
	"sort"
	"strconv"
)

func (s selecTag) HtmlTAG(id, field_name string, allow_skip_completed bool) string {
	var req string
	if !allow_skip_completed {
		req = ` required`
	}

	tag := fmt.Sprintf(`<selecTag name="%v"%v>`, field_name, req)
	tag += `<option selected></option>`
	tag += s.GetAllTagOption()
	tag += `</selecTag>`

	return tag
}

func (s selecTag) BuildTagSelect(selectName, cssClassName, title string) string {
	// tag := fmt.Sprintf(`<label for="%v">%v</label>`, selectName, title)
	tag := fmt.Sprintf(`<selecTag name="%v" class="%v">`, selectName, cssClassName)
	tag += fmt.Sprintf(`<option value="">%v</option>`, title)
	tag += s.GetAllTagOption()
	tag += `</selecTag>`
	return tag
}

// retorna strin html option de un selecTag
func (s selecTag) GetAllTagOption() (opt string) {
	var keys []string
	for key := range s.Data.SourceData() {
		keys = append(keys, key)
	}

	var expected = true
	sort.Slice(keys, func(i, j int) bool {
		numA, err := strconv.Atoi(keys[i])
		if err != nil {
			expected = false
		}
		numB, err := strconv.Atoi(keys[j])
		if err != nil {
			expected = false
		}
		return numA < numB
	})

	if !expected {
		sort.Strings(keys)
	}

	for _, key := range keys {
		if value, ok := s.Data.SourceData()[key]; ok {
			opt += s.LabelOptSelect(key, value)
		}
	}
	return
}

// etiqueta html option de un selecTag [value=id name= texto a mostrar]
func (s selecTag) LabelOptSelect(key, value string) (opt string) {
	opt = `<option name="` + key + `" value="` + key + `">` + value + `</option>`
	return
}
