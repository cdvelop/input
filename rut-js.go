package input

func (r rut) SelectedTargetChanges() string {
	if r.dni_mode {
		return "changeValueInputRunGlobal(input)"
	} else {
		return "ValidateChileanDoc(input);"
	}
}

func (r rut) InputValueChanges() string {
	if r.dni_mode {
		return "changeValueInputRunGlobal(input)"
	} else {
		return "ValidateChileanDoc(input);"
	}
}

func (r rut) JsPrivate() string {
	if r.dni_mode {

	} else {

	}

	return ""
}

func (r rut) JsListeners() string {

	if r.dni_mode {
		return `console.log("addEventListener LISTENER RUN GLOBAL");
		input_radio_type_global_run_ch.addEventListener('click',changeTypeRunGlobal,false);
		input_radio_type_global_run_ex.addEventListener('click',changeTypeRunGlobal,false);`

	} else {
		return ""
	}
}

func (r rut) JsGlobal() string {

	out := `function RunToPointFormat(rut) {
		// XX.XXX.XXX-X
		let run_number = rut.substring(0, rut.length - 2)
		let run_point = FormateaNumero(run_number);
		let _dv = rut.substring(rut.length - 2, rut.length);
	
		return run_point + _dv
	}
	const valRun = {
		// Valida el rut con su cadena completa "XXXXXXXX-X"
		validaRut: function (rutCompleto) {
			if (!/^[0-9]+[-|‐]{1}[0-9kK]{1}$/.test(rutCompleto))
				return false;
			let tmp = rutCompleto.split('-');
			let digv = tmp[1];
			let rut = tmp[0];
			if (digv == 'K') digv = 'k';
			return (valRun.dv(rut) == digv);
		},
		dv: function (T) {
			let M = 0, S = 1;
			for (; T; T = Math.floor(T / 10))
				S = (S + T % 10 * (9 - M++ % 6)) % 11;
			return S ? S - 1 : 'k';
		}
	};
	
	function ValidateChileanDoc(input) {
		// console.log("VALIDANDO RUN: ", input);
		if (valRun.validaRut(input.value)) {
			InputRight(input);
			return true
		} else {
			InputWrong(input);
			return false
		};
	};`

	out += r.jsDniMode()

	return out
}

func (r rut) jsDniMode() string {
	if r.dni_mode {
		return `
		function ValidateGlobalRunDocumentNumber(document_type,input) {
			if (input.value === "") {
				InputRight(input);
				return false
			}
			
			try {
				if (document_type === "ch") {
					return ValidateChileanDoc(input);
				} else {
					return ValidateForeignDoc(input);
				}
		
			} catch (e) {
				return ValidateChileanDoc(input);
			};
		};
		
		
		function ValidateForeignDoc(input) {
			const rex = /^[A-Za-z0-9]{5,15}$/;
			if (rex.test(input.value)) {
				InputRight(input);
				return true
			} else {
				InputWrong(input);
				return false
			}
		};	
		`
	}
	return ""
}
