	function AgeInputChange(input,module) {

		// console.log("ENTRADA EDAD: ", input);
		let age = calculateAge(input.value);
		let input_age = module.querySelector(".age-number input[type='number']");
		// console.log(" VALOR EDAD: ", age, "INPUT CHANGE AGE: ", input_age);
		if (age >= 0) {
			input_age.value = age;
		} else {
			input_age.value = '';
		}
		YearsToBirthDay(input);
		InputValidationWithValidity(input);
	};

	function YearsToBirthDay(input) {
		// console.log("YEAR TO BIRTHDAY ENTRADA : ", input);
		let input_date = module.querySelector(".age-date input[type='date']");
		let new_birthDay = ageToBirthDay(input.value, input_date.value);
		input_date.value = new_birthDay;
	};