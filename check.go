package input

import (
	"strings"

	"github.com/cdvelop/model"
)

// SourceData() map[string]string
func Check(data sourceData, onlyInternalContend bool) model.Input {
	in := check{
		Data:                  data,
		only_internal_contend: onlyInternalContend,
	}

	return model.Input{
		Component: model.Component{
			Name:        in.Name(),
			CssGlobal:   nil,
			CssPrivate:  nil,
			JsGlobal:    in,
			JsPrivate:   in,
			JsListeners: nil,
		},
		HtmlTag:  in,
		Validate: in,
		TestData: in,
	}
}

// check Box id y valor
type check struct {
	Data                  sourceData
	only_internal_contend bool
}

func (c check) Name() string {
	return c.HtmlName()
}

func (check) HtmlName() string {
	return "checkbox"
}

// validación con datos de entrada
func (c check) ValidateField(data_in string, skip_validation bool) bool {
	if !skip_validation {
		dataInArray := strings.Split(data_in, ",")
		var keysFound []string
		var badKey []string
		for _, idkeyIn := range dataInArray {
			if _, exists := (c.Data.SourceData())[idkeyIn]; exists {
				keysFound = append(keysFound, idkeyIn)
			} else {
				badKey = append(badKey, idkeyIn)
			}
		}
		if len(keysFound) > 0 && len(badKey) == 0 {
			return true
		}
	} else {
		return true
	}
	return false
}
