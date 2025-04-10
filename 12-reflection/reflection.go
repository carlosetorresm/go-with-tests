package reflection

import "reflect"

//write a function walk(x interface{}, fn func(string)) which takes a struct
// x and calls fn for all strings fields found inside. difficulty level: recursively.

func walk(x any, fn func(string)) {
	val := reflect.ValueOf(x)

	for i := range val.NumField() {
		field := val.Field(i)
		fn(field.String())
	}

}
