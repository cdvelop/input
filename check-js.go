package input

import (
	_ "embed"
)

//go:embed check.js
var check_global string

func (check) JsGlobal() string {
	return check_global
}

func (check) JsFunctions() string {
	return `crud.ModifyDomAfterCreate = Object({
		"checkbox": function (data) {
			crudFunctions.checkboxCreate(form, data);
		}
	});
	
	crud.ModifyDomAfterUpdate = Object({
		"checkbox": function (data) {
			crudFunctions.checkboxUpdate(form, data);
		}
	});
	
	crud.ModifyDomAfterDelete = Object({
		"checkbox": function (data) {
			crudFunctions.checkboxDelete(form, data);
		}
	});`
}

func (check) SelectedTargetChanges() string {
	return "TargetCheckChange(input, selected);"
}
