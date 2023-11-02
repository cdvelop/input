package input

import (
	"sort"
	"strconv"
)

func (d datalist) BuildContainerView(id, field_name string, allow_skip_completed bool) string {
	var req string
	if !allow_skip_completed {
		req = ` required`
	}

	tag := `<input list="` + field_name + `" name="` + field_name + `" id="` + id + `"` + req + ` oninput="` + DefaultValidateFunction + `">`
	tag += `<datalist id="` + id + `">`
	tag += d.GetAllTagOption()
	tag += `</datalist>`

	return tag
}

// retorna string html option de un select
func (d datalist) GetAllTagOption() (opt string) {
	var keys []string
	for key := range d.Data.SourceData() {
		keys = append(keys, key)
	}

	var result = true
	sort.Slice(keys, func(i, j int) bool {
		numA, err := strconv.Atoi(keys[i])
		if err != nil {
			result = false
		}
		numB, err := strconv.Atoi(keys[j])
		if err != nil {
			result = false
		}
		return numA < numB
	})

	if !result {
		sort.Strings(keys)
	}

	for _, key := range keys {
		if value, ok := d.Data.SourceData()[key]; ok {
			opt += d.LabelOptSelect(key, value)
		}
	}
	return
}

// etiqueta html option de un datalist
func (d datalist) LabelOptSelect(key, value string) (opt string) {
	opt = `<option data-id="` + key + `" value="` + value + `">`
	return
}
