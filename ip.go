package input

// dirección ip valida campos separados por puntos
func Ip() *ip {
	new := &ip{
		attributes: attributes{
			Title: `title="dirección ip valida campos separados por puntos ej 192.168.0.8"`,
			// Pattern: `^((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$`,
		},
		per: Permitted{
			Letters:    true,
			Numbers:    true,
			Characters: []rune{'.', ':'},
			Minimum:    7,  //IPv4 - IPv6 es 39
			Maximum:    39, // IPv6 - IPv4 es 15
		},
	}

	return new
}

type ip struct {
	attributes
	per Permitted
}

func (ip) InputName() string {
	return "Ip"
}

func (i ip) HtmlName() string {
	return "text"
}

func (i ip) BuildContainerView(id, field_name string, allow_skip_completed bool) string {
	return i.BuildHtmlTag(i.HtmlName(), "Ip", id, field_name, allow_skip_completed)

}

// validación con datos de entrada
func (i ip) ValidateField(data_in string, skip_validation bool, options ...string) (err string) {
	if !skip_validation {

		if data_in == "0.0.0.0" {
			return "ip de ejemplo no valida"
		}

		var ipV string

		if String().Contains(data_in, ":") != 0 { //IPv6
			ipV = ":"
		} else if String().Contains(data_in, ".") != 0 { //IPv4
			ipV = "."
		}

		if ipV == "" {
			return "version IPv4 o 6 no encontrada"
		}

		part := String().Split(data_in, ipV)

		if ipV == "." && len(part) != 4 {
			return "formato IPv4 no valida"
		}

		if ipV == ":" && len(part) != 8 {
			return "formato IPv6 no valida"
		}

		return i.per.Validate(data_in)
	}

	return ""
}

func (i ip) GoodTestData() (out []string) {

	out = []string{
		"120.1.3.206",
		"195.145.149.184",
		"179.183.230.16",
		"253.70.9.26",
		"215.35.117.51",
		"212.149.243.253",
		"126.158.214.250",
		"49.122.253.195",
		"53.218.195.25",
		"190.116.115.117",
		"115.186.149.240",
		"163.95.226.221",
	}

	return
}

func (i ip) WrongTestData() (out []string) {
	out = []string{
		"0.0.0.0",
		"192.168.1.1.8",
	}
	out = append(out, wrong_data...)
	return
}
