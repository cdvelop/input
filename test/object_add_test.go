package input_test

import (
	"testing"
)

func Test_Validate(t *testing.T) {
	for prueba, data := range TestData {
		t.Run((prueba), func(t *testing.T) {

			object := objects[data.Object]

			message, ok := object.ValidateData(data.ValidALL, data.Data)
			if ok != data.Expected {
				t.Fatalf("%v %v", message, ok)
			}

			// fmt.Printf("%v %v\n", message, ok)
		})
	}
}
