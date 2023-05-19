package input

//go: embed rutGlobal.js
// var rutGlobal string

func (dateAge) SelectedTargetChanges() string {
	return "AgeInputChange(input,module)"
}

func (dateAge) InputValueChanges() string {
	return "AgeInputChange(input,module)"
}
