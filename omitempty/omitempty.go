package omitempty

import (
	"encoding/json"
	"encoding/xml"
	"reflect"
	"strings"
)

type IsEmptier interface {
	IsEmpty() bool
}

func MarshalJSON(obj interface{}) ([]byte, error) {
	st := reflect.TypeOf(obj)
	fs := []reflect.StructField{}
	for i := 0; i < st.NumField(); i++ {
		fs = append(fs, st.Field(i))
	}

	for i, _ := range fs {
		if !fieldHasOmitEmpty(fs[i], "json") {
			continue
		}

		val := reflect.ValueOf(obj)
		valueField := val.Field(i)
		f := valueField.Interface()

		if isempty, ok := f.(IsEmptier); ok {
			if !isempty.IsEmpty() {
				continue
			}
			fs[i].Tag = reflect.StructTag(`json:"-"`)
		}
		continue
	}

	st2 := reflect.StructOf(fs)

	v := reflect.ValueOf(obj)
	v2 := v.Convert(st2)
	return json.Marshal(v2.Interface())
}

func MarshalXML(obj interface{}, e *xml.Encoder, start xml.StartElement) error {
	st := reflect.TypeOf(obj)
	fs := []reflect.StructField{}
	for i := 0; i < st.NumField(); i++ {
		fs = append(fs, st.Field(i))
	}

	for i, _ := range fs {
		if !fieldHasOmitEmpty(fs[i], "xml") {
			continue
		}

		val := reflect.ValueOf(obj)
		valueField := val.Field(i)
		f := valueField.Interface()

		if isempty, ok := f.(IsEmptier); ok {
			if !isempty.IsEmpty() {
				continue
			}
			fs[i].Tag = reflect.StructTag(`xml:"-"`)
		}
		continue
	}

	st2 := reflect.StructOf(fs)

	v := reflect.ValueOf(obj)
	v2 := v.Convert(st2)
	return e.EncodeElement(v2.Interface(), start)
}

func fieldHasOmitEmpty(field reflect.StructField, encoder string) bool {
	// Get the field tag value
	tag := field.Tag.Get(encoder)
	options := strings.Split(tag, ",")
	for _, option := range options {
		if option == "omitempty" {
			return true
		}
	}
	return false
}
