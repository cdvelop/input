package input

import (
	"strconv"

	"github.com/cdvelop/model"
)

// options:
// ej: min="2", max="10", hidden....
// min mínimo de caracteres permitidos ej: 3 o 5 ... min default 5
// max máximo de caracteres permitidos ej: 20 50 ... max default 50
// Pattern_start="^[A-Za-zÑñ 0-9:.-]{"
// Pattern_end="}$"
func Password(options ...string) model.Input {
	in := password{
		attributes: attributes{},
		Permitted: Permitted{
			Letters:    true,
			Tilde:      true,
			Numbers:    true,
			Characters: []rune{' ', '$', '#', '%', '?', '.', ',', '-', '_'},
			Minimum:    5,
			Maximum:    50,
		},
	}
	in.Set(options...)

	if in.Min != "" {
		in.Minimum, _ = strconv.Atoi(in.Min)
	}

	if in.Max != "" {
		in.Maximum, _ = strconv.Atoi(in.Max)
	}

	return model.Input{
		InputName: in.Name(),
		Tag:       &in,
		Validate:  &in,
		TestData:  &in,
	}
}

// Solo letras (en cualquier caso), números, guiones, guiones bajos y puntos.
// (No el carácter de barra, que se usa para escapar del punto).
type password struct {
	attributes
	Permitted
}

func (p password) Name() string {
	return p.HtmlName()
}

func (p password) HtmlName() string {
	return "password"
}

func (p password) HtmlTag(id, field_name string, allow_skip_completed bool) string {
	return p.BuildHtmlTag(p.HtmlName(), p.Name(), id, field_name, allow_skip_completed)
}

func (p password) GoodTestData() (out []string) {

	temp := []string{
		"c0ntra3",
		"M1 contraseÑ4",
		"contrase",
		"cont",
		"12345",
		"UNA Frase tambien Cuenta",
		"DOS Frases tambien CuentaN",
		"CUATRO FraseS tambien CuentaN",
	}

	for _, v := range temp {
		if len(v) >= p.Minimum && len(v) <= p.Maximum {
			out = append(out, v)
		}
	}

	return
}

func (p password) WrongTestData() (out []string) {

	temp := []string{
		"",
		"Ñ",
		"c",
		" ",
		"2",
		"%",
		"sdlksññs092830928309280%%%%%9382¿323294720&&/0kdlskdlskdskdñskdlskdsññdkslkdñskdslkdsñ",
		"sdlksññs0928309283092809382%%¿323294720&&/0kdlskdlskdskdñskdlskdsññdkslkdñskdslkdsñ",
		"sdlksññs0928309283092809382¿78%%323294720&&/0kdlskdlskdskdñskdlskdsññdkslkdñskdslkdsñ",
	}

	return temp
}
