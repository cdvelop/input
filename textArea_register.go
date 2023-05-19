package input

import (
	"bytes"
	_ "embed"
	"log"
	"text/template"
)

//go:embed textArea_global.js
var global_text_area string

func (textArea) SelectedTargetChanges() string {
	return "TextAreaAutoGrow(input);TextAreaValidate(input);"
}
func (textArea) InputValueChanges() string {
	return "TextAreaAutoGrow(input);TextAreaValidate(input);"
}

func (t textArea) updatePatternInJs() {
	p, err := template.New("").Parse(global_text_area)
	if err != nil {
		log.Println(err)
		return
	}

	var html bytes.Buffer
	err = p.Execute(&html, t.pattern)
	if err != nil {
		log.Println(err)
		return
	}

	global_text_area = html.String()

	// fmt.Println(global_text_area)
}
