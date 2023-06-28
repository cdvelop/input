package input_test

import (
	"github.com/cdvelop/input"
	"github.com/cdvelop/model"
)

const (
	TableName1 = "usuario"
	TableName2 = "shopping_cart"

	id_user_key    = "id_usuario"
	nameKey        = "name"
	genderKey      = "gender"
	descriptionKey = "description"
	rutKey         = "rut_usuario"
)

var (
	objects = map[string]model.Object{
		Object1().Name: *Object1(),
		Object2().Name: *Object2(),
	}
)

type dataGenero struct{}

func (dataGenero) SourceData() map[string]string {
	return map[string]string{"D": "Dama", "V": "Varón"}
}

func Object1() *model.Object {
	t1 := model.Object{
		ApiHandler: model.ApiHandler{
			Name: TableName1,
		},
		Fields: []model.Field{
			{Name: id_user_key, Input: input.Pk()},
			{Name: nameKey, Input: input.Text(), Legend: "Nombre"},
			{Name: genderKey, Input: input.Radio(dataGenero{})},
			{Name: descriptionKey, Input: input.Text(), Legend: "Descripción", SkipCompletionAllowed: true, SkipValidation: true},
			{Name: rutKey, Input: input.Rut(), Inalterable: true, SkipCompletionAllowed: true},
		},
	}

	return &t1
}

type dataState struct{}

func (dataState) SourceData() map[string]string {
	return map[string]string{
		"inuse": "En Uso", "lost": "Perdido",
	}
}

func Object2() *model.Object {

	t2 := model.Object{
		ApiHandler: model.ApiHandler{
			Name: TableName2,
		},
		Fields: []model.Field{
			{Name: "id_" + TableName2, Input: input.Pk()},
			{Name: id_user_key, Legend: "Id Usuario", Input: input.Pk()},
			{Name: TableName2 + "_state", Input: input.Radio(dataState{})},
		},
	}

	return &t2
}
