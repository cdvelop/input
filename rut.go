package input

type rut struct {
	name        string
	hide_typing bool
	dni_mode    bool
	attributes
	dni Permitted
}

// parámetro opcionales:
// hide-typing: ocultar información  al escribir
// dni-mode: acepta documentos extranjeros
func Rut(options ...string) *rut {
	new := rut{
		name: "Rut",
		attributes: attributes{
			Autocomplete: `autocomplete="off"`,
			Class:        `class="rut"`,
			DataSet:      `data-option="ch"`,
		},
		dni: Permitted{
			Letters: true,
			Numbers: true,
			Minimum: 9,
			Maximum: 15,
		},
	}

	for _, opt := range options {
		switch opt {
		case "hide-typing":
			new.hide_typing = true
		case "dni-mode":
			new.dni_mode = true
		}
	}

	if new.dni_mode {
		new.name = "RutDni"
		new.Title = `title="Documento Chileno (ch) o Extranjero (ex)"`
		if !new.hide_typing {
			new.PlaceHolder = `placeholder="ej: (ch) 11222333-k  /  (ex) 1b2334"`
		}

		// new.Pattern = `^[A-Za-z0-9]{9,15}$`
		new.Maxlength = `maxlength="15"`
	} else {
		new.Title = `title="rut sin puntos y con guion ejem.: 11222333-4"`
		if !new.hide_typing {
			new.PlaceHolder = `placeholder="ej: 11222333-4"`
		}
		new.Maxlength = `maxlength="10"`

		// new.Pattern = `^[0-9]+-[0-9kK]{1}$`
	}

	return &new
}

func (r rut) InputName() string {
	return r.name
}

func (r rut) HtmlName() string {
	if r.hide_typing {
		return "password"
	}
	return "text"
}

func (r rut) BuildContainerView(id, field_name string, allow_skip_completed bool) string {

	if r.dni_mode {

		tag := `<div class="run-type">`

		tag += r.BuildHtmlTag(r.HtmlName(), r.name, id, field_name, allow_skip_completed)

		tag += `<div class="rut-label-container"><label class="rut-radio-label">
			<input type="radio" name="type-dni" data-name="dni-ch" value="ch" checked="checked" onchange="changeDniType(this)">
			<span title="Documento Chileno">ch</span>
		</label>
	
		<label class="rut-radio-label">
			<input type="radio" name="type-dni" data-name="dni-ex" value="ex" onchange="changeDniType(this)">
			<span title="Documento Extranjero">ex</span>
		</label>
	  </div>
    </div>`

		return tag

	} else {
		return r.BuildHtmlTag(r.HtmlName(), r.name, id, field_name, allow_skip_completed)
	}
}
