package input

import "github.com/cdvelop/model"

// options: title="xxx"
func DayWord(options ...string) model.Input {
	in := dayWord{
		attributes: attributes{},
	}
	in.Set(options...)

	return model.Input{
		Component: model.Component{
			Name:        in.Name(),
			CssGlobal:   nil,
			CssPrivate:  nil,
			JsGlobal:    nil,
			JsPrivate:   nil,
			JsListeners: nil,
		},
		Build:    in,
		Validate: in,
		TestData: in,
	}
}

// formato dia DD como palabra ej. Lunes 24 Diciembre
type dayWord struct {
	attributes
}

func (d dayWord) Name() string {
	return "dayword"
}

func (d dayWord) HtmlName() string {
	return "text"
}

func (dayWord) GoodTestData(table_name, field_name string, random bool) (out []string) {
	return MonthDay().GoodTestData("", "", random)
}

func (dayWord) WrongTestData() (out []string) {
	return MonthDay().WrongTestData()
}
