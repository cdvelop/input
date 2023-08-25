function CheckChange(input) {
	const father = input.parentNode.parentNode
	const checkboxes = father.querySelectorAll('input[type="checkbox"]');
	let all_selected = true;

	for (const cb of checkboxes) {
		if (cb.checked) {
			all_selected = false;
			break;
		}
	}

	if (all_selected) {
		if (input.hasAttribute('required')) {
			InputWrong(input);
		} else {
			InputNormal(input);
		}
		// console.log("Ningún checkbox está seleccionado.");
	} else {
		InputRight(input);
		// console.log("Al menos un checkbox está seleccionado.");
	}
}

function checkboxCreate(form, data) {
	const check_container = form.querySelector('[data-name="' + data.name + '"] label[data-id="' + data.id + '"]');
	if (check_container != null) {
		check_container.insertAdjacentHTML("beforeend", data.tag);
	}
};

function checkboxUpdate(form, data) {
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

function checkboxDelete(form, data) {
	let input_check = form.querySelector('[data-name="' + data.name + '"] label[data-id="' + data.id + '"]');
	if (input_check != null) {
		input_check.parentNode.removeChild(input_check);
	}
};