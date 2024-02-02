package input

// add ej: `min="` only send: "min" new delete
func (a *attributes) add(option, delete string) string {
	out := String().Replace(option, delete+`="`, "")
	out = String().TrimSuffix(out, `"`)
	return out
}

func (a *attributes) Set(options ...string) {

	for _, option := range options {
		switch {

		case String().Contains(option, "min=") != 0:
			a.Min = a.add(option, "min")

		case String().Contains(option, "max=") != 0:
			a.Max = a.add(option, "max")

		case String().Contains(option, "maxlength=") != 0:
			a.Maxlength = option

		case String().Contains(option, "data-") != 0:
			a.DataSet = option

		case String().Contains(option, "class=") != 0:
			a.Class = option

		case String().Contains(option, "placeholder=") != 0:
			a.PlaceHolder = option

		case String().Contains(option, "title=") != 0:
			a.Title = option

		case String().Contains(option, "autocomplete=") != 0:
			a.Autocomplete = option

		case String().Contains(option, "rows=") != 0:
			a.Rows = option
		case String().Contains(option, "cols=") != 0:
			a.Cols = option

		case String().Contains(option, "step=") != 0:
			a.Step = option

		case String().Contains(option, "oninput=") != 0:
			a.Oninput = option

		case String().Contains(option, "onkeyup=") != 0:
			a.Onkeyup = option

		case String().Contains(option, "onchange=") != 0:
			a.Onchange = option

		case String().Contains(option, "accept=") != 0:
			a.Accept = option

		case option == "multiple":
			a.Multiple = option

		}
	}

}
