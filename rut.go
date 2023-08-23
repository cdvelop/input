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
			Class:        `class="rut"`,
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
		in.PlaceHolder = `placeholder="ej: (ch) 11222333-k  /  (ex) 1b2334"`
		in.Pattern = `^[A-Za-z0-9]{9,15}$`
		in.Maxlength = `maxlength="15"`
	} else {
		in.Title = `title="rut sin puntos y con guion ejem.: 11222333-4"`
		in.PlaceHolder = `placeholder="ej: 11222333-4"`
		in.Maxlength = `maxlength="10"`
		// in.Pattern = `^[0-9]+-[0-9kK]{1}$`
	}

	return model.Input{
		InputName: in.Name(),
		Tag:       &in,
		Validate:  &in,
		TestData:  &in,
	}
}

func (r rut) Name() string {
	if r.dni_mode {
		return "rut_dni"
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

		tag += `<div class="rut-label-container"><label class="rut-radio-label">
			<input type="radio" name="type-dni" data-name="dni-ch" value="ch" checked="checked" oninput="changeDniType(this, this.form)">
			<span title="Documento Chileno">ch</span>
		</label>
	
		<label class="rut-radio-label">
			<input type="radio" name="type-dni" data-name="dni-ex" value="ex" oninput="changeDniType(this, this.form)">
			<span title="Documento Extranjero">ex</span>
		</label>
	  </div>
    </div>`

		return tag

	} else {
		return r.BuildHtmlTag(r.HtmlName(), r.Name(), id, field_name, allow_skip_completed)
	}
}
