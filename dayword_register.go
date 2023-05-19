package input

// go:embed js/global.js
// var js_global string

// go:embed css/style.css
// var css_style string

func init() {
	// d := dayWord{}

	// devui.RegisterCSS(d.Name(), "css", nil, css_style)

	// js_handler.RegisterJsGlobalFunctions(d.Name(), nil, js_global)

	// form_handler.RegisterInputsJsModule(d.Name(), nil, nil, d, d)
}

func (dayWord) SelectedTargetChanges() string {
	return "DayWordChange(input);"
}

func (dayWord) InputValueChanges() string {
	return "DayWordChange(input);"
}
