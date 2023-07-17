package input

import (
	"github.com/cdvelop/model"
)

type rut struct {
	hide_typing bool
	dni_mode    bool
	attributes
}

// parámetro opcionales:
// hide-typing: ocultar información  al escribir
// dni-mode: acepta documentos extranjeros
func Rut(options ...string) model.Input {
	in := rut{
		attributes: attributes{
			Autocomplete: `autocomplete="off"`,
		},
	}

	for _, opt := range options {
		switch opt {
		case "hide-typing":
			in.hide_typing = true
		case "dni-mode":
			in.dni_mode = true
		}
	}

	if in.dni_mode {
		in.Title = `title="Documento Chileno (ch) o Extranjero (ex)"`
		in.PlaceHolder = `placeholder="ej: chileno (ch) 11222333-k / extranjero sin dígito ni guion (ex) 1b2334"`
		in.Pattern = `^[A-Za-z0-9]{9,15}$`
		in.Maxlength = `maxlength="15"`
	} else {
		in.Title = `title="rut sin puntos y con guion ejem.: 11222333-4"`
		in.PlaceHolder = `placeholder="ej: 11222333-4"`
		in.Maxlength = `maxlength="10"`
		// in.Pattern = `^[0-9]+-[0-9kK]{1}$`
	}

	return model.Input{
		Name:     in.Name(),
		Tag:      in,
		Validate: in,
		TestData: in,
	}
}

func (r rut) Name() string {
	if r.dni_mode {
		return "dni"
	}

	return "rut"
}

func (r rut) HtmlName() string {
	if r.hide_typing {
		return "password"
	}
	return "text"
}

func (r rut) HtmlTag(id, field_name string, allow_skip_completed bool) string {

	if r.dni_mode {

		tag := `<div class="run-type">`

		tag += r.BuildHtmlTag(r.HtmlName(), r.Name(), id, field_name, allow_skip_completed)

		tag += `<label class="block-label">
			<input type="radio" name="type-dni" data-name="dni-ch" value="ch" checked="checked" oninput="changeDniType(this)">
			<span title="Documento Chileno">ch</span>
		</label>
	
		<label class="block-label">
			<input type="radio" name="type-dni" data-name="dni-ex" value="ex" oninput="changeDniType(this)">
			<span title="Documento Extranjero">ex</span>
		</label>
	</div>`

		return tag

	} else {
		return r.BuildHtmlTag(r.HtmlName(), r.Name(), id, field_name, allow_skip_completed)
	}
}

func (r rut) Css() string {
	return `.run-type{
    display: flex;
    flex-direction: row;
    max-width: 25vw;
}`
}

func (r rut) JsGlobal() string {

	if r.dni_mode {
		return `function changeDniType(e) {
			const input_dni = e.closest('.run-type').querySelector('input[type="text"][data-name="dni"]');
			if (e.value === "ch") {
				input_dni.setAttribute("maxlength", 10);
			} else {
				input_dni.setAttribute("maxlength", 15);
			}
			validateField(input_dni);
		}`
	}

	return `function RunToPointFormat(rut) {
		// XX.XXX.XXX-X
		let run_number = rut.substring(0, rut.length - 2)
		let run_point = FormateaNumero(run_number);
		let _dv = rut.substring(rut.length - 2, rut.length);
	
		return run_point + _dv
	}`
}
