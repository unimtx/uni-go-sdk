package uni

import (
	"reflect"
	"strings"
)

func IsEmpty(obj interface{}) bool {
	if obj == nil {
		return true
	}

	val := reflect.ValueOf(obj)

	switch val.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice:
		return val.Len() == 0
	case reflect.Ptr:
		if val.IsNil() {
			return true
		}
		ref := val.Elem().Interface()
		return IsEmpty(ref)
	default:
		zero := reflect.Zero(val.Type())
		return reflect.DeepEqual(obj, zero.Interface())
	}
}

func ToLowerFirstChar(s string) string {
	return strings.ToLower(string(s[0])) + s[1:]
}

func StructToMap(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	data := make(map[string]interface{})

	for i := 0; i < t.NumField(); i++ {
		key := ToLowerFirstChar(t.Field(i).Name)
		val := v.Field(i).Interface()

		if (IsEmpty(val)) {
			continue
		}
		data[key] = val
	}
	return data
}
