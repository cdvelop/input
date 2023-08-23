
function TargetCheckChange(inputs, data_selected) {

	console.log("INPUTS CHECK: ",inputs," DATA: ",data_selected)
	let counter = 0;
	let targetValues = data_selected[i].dataset.value.split(",");
	for (let i = 0; i < targetValues.length; i++) {
		for (let c = 0; c < inputs.length; c++) {
			// console.log("VALOR FORM ",input[c].value, "VALOR TARGET", targetValues[i]);
			if (inputs[c].value == targetValues[i]) {
				inputs[c].checked = true;
				counter++;
				break;
			}
		}
	}

	if (counter > 0) {
		changeFieldsetColor(inputs[0], true);
	} else {
		changeFieldsetColor(inputs[0], false);
	}
}

crudFunctions.checkboxCreate = function (form, data) {
	const check_container = form.querySelector('[data-name="' + data.name + '"] label[data-id="' + data.id + '"]');
	if (check_container != null) {
		check_container.insertAdjacentHTML("beforeend", data.tag);
	}
};

crudFunctions.checkboxUpdate = function (form, data) {
	let label_check = form.querySelector('[data-name="' + data.name + '"] label[data-id="' + data.id + '"]');
	if (label_check != null) {
		label_check.innerHTML = data.tag;

		let input_check = label_check.querySelector('input[type=checkbox]');
		let new_for_id = module.id + '.' + data.name + '.' + data.id;

		label_check.setAttribute('for', new_for_id);
		input_check.id = new_for_id;
		label_check.dataset.id = data.id;
	}
};

crudFunctions.checkboxDelete = function (form, data) {
	let input_check = form.querySelector('[data-name="' + data.name + '"] label[data-id="' + data.id + '"]');
	if (input_check != null) {
	  input_check.parentNode.removeChild(input_check);
	}
  };