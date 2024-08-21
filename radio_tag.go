package input

import (
	"fmt"
	"sort"
	"strconv"
)

func (r radio) BuildInputHtml(id, fieldName string) string {
	var id3 string

	keys := make([]string, 0, len(r.Data.SourceData()))
	for k := range r.Data.SourceData() {
		keys = append(keys, k)
	}
	// sort.Strings(keys)

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

	var tags string
	for i, value := range keys {
		id3 = fmt.Sprintf("%v.%v", id, i)
		tags += `<label for="` + id3 + `" class="block-label">`

		r.Value = `value="` + value + `"`

		tags += r.BuildHtmlTag("radio", r.name, id3, fieldName)

		tags += `<span>` + r.Data.SourceData()[value] + `</span>`
		tags += `</label>`
	}
	return tags
}
