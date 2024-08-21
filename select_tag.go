package input

import (
	"fmt"
	"sort"
	"strconv"
)

func (s selecTag) BuildInputHtml(id, fieldName string) string {
	var req string
	if !s.AllowSkipCompleted {
		req = ` required`
	}

	tag := fmt.Sprintf(`<select name="%v" oninput="`+DefaultValidateFunction+`"%v>`, fieldName, req)
	tag += `<option selected></option>`
	tag += s.GetAllTagOption()
	tag += `</select>`

	return tag
}

// retorna string html option de un selecTag
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
