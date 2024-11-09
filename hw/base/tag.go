package base

import "reflect"

func getStructTag(f reflect.StructField) string {
	return string(f.Tag)
}
