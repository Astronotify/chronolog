package internal

import (
	"reflect"
	"runtime/debug"
)

func MergeAdditionalData(data ...map[string]any) map[string]any {
	merged := map[string]any{}
	for _, d := range data {
		for k, v := range d {
			merged[k] = v
		}
	}
	return merged
}

func ClassifyError(err error) string {
	return "GenericError"
}

func ExtractStackTrace(err error) string {
	return string(debug.Stack())
}

func GetStructName(i interface{}) string {
	t := reflect.TypeOf(i)
	if t.Kind() == reflect.Pointer {
		t = t.Elem()
	}
	return t.Name()
}
