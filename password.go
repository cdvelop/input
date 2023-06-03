package input

import (
	"log"
	"regexp"
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
	}
	in.Set(options...)

	//  si no me envían pattern
	if in.Pattern == "" {
		if in.Min == "" {
			in.Min = "5"
		}

		if in.Max == "" {
			in.Max = "50"
		}
		in.Pattern_start = `^[A-Za-zÑñ 0-9:.-]{`
		in.Pattern_end = `}$`
		in.patternUpdate()
	}

	return model.Input{
		Component: model.Component{
			Name:        in.Name(),
			CssGlobal:   nil,
			CssPrivate:  nil,
			JsGlobal:    nil,
			JsPrivate:   nil,
			JsListeners: nil,
		},
		Tag:      in,
		Validate: in,
		TestData: in,
	}
}

// Solo letras (en cualquier caso), números, guiones, guiones bajos y puntos.
// (No el carácter de barra, que se usa para escapar del punto).
type password struct {
	attributes
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

// validación con datos de entrada
func (p password) ValidateField(data_in string, skip_validation bool) (ok bool) {
	if !skip_validation {

		pvalid := regexp.MustCompile(p.Pattern)

		return pvalid.MatchString(data_in)

	} else {
		return true
	}
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

	if p.Min != "" && p.Max != "" {

		min, err := strconv.Atoi(p.Min)
		if err != nil {
			// Manejar el error si la conversión falla
			log.Fatal("No se pudo convertir el string min: " + p.Min + " a int")
			return
		}

		max, err := strconv.Atoi(p.Max)
		if err != nil {
			// Manejar el error si la conversión falla
			log.Fatal("No se pudo convertir el string max: " + p.Max + " a int")
			return
		}

		if min > 0 {

			for _, pwd := range temp {
				if len(pwd) >= min && len(pwd) <= max {
					out = append(out, pwd)
				}
			}

			return
		}
	}

	return temp
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

	if p.Min != "" && p.Max != "" {

		min, err := strconv.Atoi(p.Min)
		if err != nil {
			// Manejar el error si la conversión falla
			log.Fatal("No se pudo convertir el string min: " + p.Min + " a int")
			return
		}

		max, err := strconv.Atoi(p.Max)
		if err != nil {
			// Manejar el error si la conversión falla
			log.Fatal("No se pudo convertir el string max: " + p.Max + " a int")
			return
		}

		if min > 0 {

			for _, pwd := range temp {
				if len(pwd) != min || len(pwd) > max {
					out = append(out, pwd)
				}
			}

			return
		}

	}

	return temp
}
