function RadioTargetChange(input,selected) {
    for (let r = 0; r < input.length; r++) {
		if (input[r].value == selected[i].dataset.value) {
			input[r].checked = true;
			InputValidationWithValidity(input[r]);
			break;
		}
	}
}