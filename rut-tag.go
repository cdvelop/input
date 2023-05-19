package input

func (r rut) HtmlTAG(id, field_name string, allow_skip_completed bool) string {

	if r.dni_mode {

		var required string
		if !allow_skip_completed {
			required = ` required`
		}

		return `<div class="run-type">
		<input type="` + r.HtmlName() + `" maxlength="12" name="` + field_name + `" data-name="global_run" title="` + r.Title + `"
			pattern="` + r.Pattern + `" placeholder="` + r.PlaceHolder + `"` + required + `>
		<label class="block-label">
			<input type="radio" name="type-global-run" data-name="global-run-ch" value="ch" checked="checked" do-not-send-with-form>
			<span title="Documento Chileno">ch</span>
		</label>
	
		<label class="block-label">
			<input type="radio" name="type-global-run" data-name="global-run-ex" value="ex" do-not-send-with-form>
			<span title="Documento Extranjero">ex</span>
		</label>
	</div>`

	} else {

		return r.Build(r.HtmlName(), r.Name(), id, field_name, allow_skip_completed)

	}

}
