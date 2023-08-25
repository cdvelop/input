package input

import (
	"fmt"
	"sort"
	"strconv"
)

func (c check) HtmlTag(id, field_name string, allow_skip_completed bool) string {

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
		id3 := fmt.Sprintf("%v.%v", id, i)

		tags += c.newTag(id3, field_name, value, (c.Data.SourceData())[value], c.only_internal_contend)
	}

	return tags
}

func (c check) newTag(id, field_name, field_value, text_field string, only_internal_contend bool) string {

	if id == "" {
		id = field_value
	}

	tag_input := fmt.Sprintf(`<input type="checkbox" id="%v" name="%v" value="%v" onchange="CheckChange(this)"><span>%v</span>`,
		id, field_name, field_value, text_field)

	if !c.only_internal_contend {
		return fmt.Sprintf(`<label data-id="%v" for="%v" class="block-label">%v</label>`, field_value, id, tag_input)
	} else {
		return tag_input
	}
}
