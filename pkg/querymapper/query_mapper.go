package querymapper

import (
	"reflect"
	"strconv"
)

func GetArrayOfPropertiesFrom(s interface{}) []interface{} {
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)
	props := getPropertiesMap(t, v)
	arr := getArrayFromProperties(props)
	return arr
}

func getPropertiesMap(t reflect.Type, v reflect.Value) map[int]interface{} {
	props := make(map[int]interface{})
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		posStr := field.Tag.Get("position")
		pos, err := strconv.Atoi(posStr)
		if err != nil {
			panic(err)
		}
		props[pos] = v.Field(i).Interface()
	}
	return props
}

func getArrayFromProperties(props map[int]interface{}) []interface{} {
	arr := make([]interface{}, len(props))
	for i := 0; i < len(props); i++ {
		arr[i] = props[i]
	}
	return arr
}
