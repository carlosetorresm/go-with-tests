package reflection

import "reflect"

//write a function walk(x interface{}, fn func(string)) which takes a struct
// x and calls fn for all strings fields found inside. difficulty level: recursively.

func walk(x any, fn func(string)) {
	val := getValue(x)

	numberOfValues := 0
	var getField func(int) reflect.Value

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		numberOfValues = val.NumField()
		getField = val.Field
	case reflect.Slice, reflect.Array:
		numberOfValues = val.Len()
		getField = val.Index
	}

	for i := range numberOfValues {
		walk(getField(i).Interface(), fn)
	}
}

func getValue(x any) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}
	return val
}
