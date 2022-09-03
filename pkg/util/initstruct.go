package util

import (
	"reflect"
)

func InitStruct[T any](i T) T {
	v := reflect.ValueOf(&i).Elem()
	return reflect.New(v.Type().Elem()).Interface().(T)
}
