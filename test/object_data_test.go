package input_test

type kv map[string]string

var (
	TestData = map[string]struct {
		Object   string
		Data     map[string]string
		ValidALL bool //validar como nuevo (todo) o solo campos requeridos presentes
		Expected bool //resultado esperado
	}{
		"todos los campos correctos?": {TableName1,
			kv{nameKey: "Luis", genderKey: "D", descriptionKey: "no tiene", rutKey: "73528171-2"},
			true, true},

		"tabla existe?": {"tablex",
			kv{nameKey: "Marco", genderKey: "D", descriptionKey: "no tiene"},
			true, false},

		"todos los campos, nombre correcto?": {TableName1,
			kv{nameKey: "Luis#1", genderKey: "D", descriptionKey: "no tiene"},
			true, false},

		"solo campos requeridos?": {TableName1,
			kv{nameKey: "Maria Joaquina", genderKey: "D"},
			true, true},

		"id + todos los campos correctos?": {TableName1,
			kv{id_user_key: "123456789", nameKey: "Luis", genderKey: "D", descriptionKey: "no tiene"},
			true, true},

		"new todos los campos?": {TableName1,
			kv{nameKey: "Luis"},
			true, false},

		"nombre solo texto correcto?": {TableName1,
			kv{nameKey: "1u50"},
			false, false},

		"llave primaria valida en update?": {TableName1,
			kv{id_user_key: "OR"},
			false, false},

		"llave primaria existente en update?": {TableName1,
			kv{genderKey: "D"},
			false, false},

		// test en tabla 2
		"campos correctos? tabla 2 ": {TableName2,
			kv{
				id_user_key:           "222",
				TableName2 + "_state": "inuse",
				nameKey:               "<script>",
				genderKey:             "=",
			},
			true, false},

		"todos los campos pertenecen a tabla 2?": {TableName2,
			kv{
				id_user_key:           "222",
				TableName2 + "_state": "inuse",
				nameKey:               "Juana",
			},
			true, false},

		"campos correctos tabla 2?": {TableName2,
			kv{
				id_user_key:           "222",
				TableName2 + "_state": "inuse",
			},
			true, true},

		"tabla 2 nueva ok sin id ni fk id?": {TableName2,
			kv{
				TableName2 + "_state": "lost",
			},
			true, false},

		"modificar campo inalterable?": {TableName1,
			kv{id_user_key: "123456789", rutKey: "12.03"},
			false, false},
	}
)
