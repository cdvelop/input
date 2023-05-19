package input

import (
	_ "embed"
)

// go:embed global_radio.js
// var global_radio string

func init() {
	// r := radio{}

	// js_handler.RegisterJsGlobalFunctions(r.Name(), nil, global_radio)

	// form_handler.RegisterInputsJsModule(r.Name(), nil, nil, r, nil)
}

func (radio) SelectedTargetChanges() string {
	return "RadioTargetChange(input,selected);"
}
