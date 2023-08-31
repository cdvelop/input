package input

import "strings"

// add ej: `min="` only send: "min" in delete
func (a *attributes) add(option, delete string) string {
	out := strings.Replace(option, delete+`="`, "", 1)
	out = strings.TrimSuffix(out, `"`)
	return out
}

func (a *attributes) Set(options ...string) {

	for _, option := range options {
		switch {

		case strings.Contains(option, "min="):
			a.Min = a.add(option, "min")

		case strings.Contains(option, "max="):
			a.Max = a.add(option, "max")

		case strings.Contains(option, "maxlength="):
			a.Maxlength = option

		case strings.Contains(option, "data-"):
			a.DataSet = option

		case strings.Contains(option, "class="):
			a.Class = option

		case strings.Contains(option, "placeholder="):
			a.PlaceHolder = option

		case strings.Contains(option, "title="):
			a.Title = option

		case strings.Contains(option, "autocomplete="):
			a.Autocomplete = option

		case strings.Contains(option, "rows="):
			a.Rows = option
		case strings.Contains(option, "cols="):
			a.Cols = option

		case strings.Contains(option, "step="):
			a.Step = option

		case strings.Contains(option, "oninput="):
			a.Oninput = option

		case strings.Contains(option, "onkeyup="):
			a.Onkeyup = option

		case strings.Contains(option, "onchange="):
			a.Onchange = option

		case strings.Contains(option, "accept="):
			a.Accept = option

		case option == "multiple":
			a.Multiple = option

		}
	}

}
