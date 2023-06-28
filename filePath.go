package input

import (
	"regexp"

	"github.com/cdvelop/model"
)

// options:
// "multiple"
// accept="image/*"
// title="Imágenes jpg"
func FilePath(options ...string) model.Input {
	in := filePath{
		attributes: attributes{
			Pattern: `^(\.\/|\.\?\\|\/)?([\w\s.-]+[\\\/]?)*$`,
		},
	}
	in.Set(options...)

	return model.Input{
		Object: model.Object{
			ApiHandler: model.ApiHandler{
				Name: in.Name(),
			},
			Css:         nil,
			JsGlobal:    nil,
			JsFunctions: nil,
			JsListeners: nil,
		},
		Tag:      in,
		Validate: in,
		TestData: in,
	}
}

type filePath struct {
	attributes
}

func (f filePath) Name() string {
	return "filePath"
}

func (f filePath) HtmlName() string {
	return "file"
}

func (f filePath) HtmlTag(id, field_name string, allow_skip_completed bool) (tags string) {
	return f.BuildHtmlTag(f.HtmlName(), f.Name(), id, field_name, allow_skip_completed)
}

// validación con datos de entrada
func (f filePath) ValidateField(data_in string, skip_validation bool) (ok bool) {
	if !skip_validation {

		if len(data_in) == 2 {
			return false
		}

		pvalid := regexp.MustCompile(f.Pattern)

		return pvalid.MatchString(data_in)
	} else {
		return true
	}
}

func (f filePath) GoodTestData() (out []string) {

	return []string{
		".\\misArchivos",
		".\\todos\\videos",
	}
}

func (f filePath) WrongTestData() (out []string) {
	out = []string{
		"\\-",
		"///.",
	}
	return
}
