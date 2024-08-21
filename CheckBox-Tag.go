package input

import (
	"sort"
	"strconv"
)

func (c check) BuildInputHtml(id, fieldName string) string {

	keys := make([]string, 0, len(c.Data.SourceData()))
	for k := range c.Data.SourceData() {
		keys = append(keys, k)
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
	var tags string
	for i, value := range keys {
		id3 := id + "." + strconv.Itoa(i)

		tags += c.newTag(id3, fieldName, value, (c.Data.SourceData())[value])
	}

	return tags
}

func (c check) newTag(id, fieldName, field_value, text_field string) string {

	if id == "" {
		id = field_value
	}

	tag_input := `<input type="checkbox" id="` + id + `" name="` + fieldName + `" value="` + field_value + `" onchange="CheckChange(this)"><span>` + text_field + `</span>`

	if !c.onlyInternalContend {
		return `<label data-id="` + field_value + `" for="` + id + `" class="block-label">` + tag_input + `</label>`
	} else {
		return tag_input
	}
}
