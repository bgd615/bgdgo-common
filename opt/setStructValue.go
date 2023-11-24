package opt

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

func SetValueFromJsonTag(src interface{}, jsonTag string, value interface{}) error {
	valueOf := reflect.ValueOf(src)
	kind := valueOf.Kind()
	if kind == reflect.Ptr {
		valueOf = valueOf.Elem()
		kind = valueOf.Kind()
	}
	if kind == reflect.Struct {
		return setStructValue(valueOf, jsonTag, value)
	}
	if kind == reflect.Slice || kind == reflect.Array {
		return setSliceValue(valueOf, jsonTag, value)
	}
	return errors.New("only struct type are supported")
}
func setStructValue(valueOf reflect.Value, jsonTag string, value interface{}) (err error) {
	defer func() {
		if rerr := recover(); rerr != nil {
			err = errors.New(fmt.Sprint(rerr))
		}
	}()
	vOf := reflect.ValueOf(value)
	if vOf.Kind() == reflect.Ptr {
		vOf = vOf.Elem()
	}
	for i := 0; i < valueOf.NumField(); i++ {
		typeField := valueOf.Type().Field(i)
		valueField := valueOf.Field(i)
		//
		if typeField.Anonymous {
			if err = setStructValue(valueField, jsonTag, value); err == nil {
				return nil
			}
		}

		if getJsonTag(typeField) == jsonTag {
			vu, err := getNewValue(valueField, vOf, jsonTag)
			if err != nil {
				return err
			}
			valueField.Set(vu)
			return nil
		}
	}
	return errors.New(fmt.Sprintf("object does not have the json:%v", jsonTag))
}
func setSliceValue(valueOf reflect.Value, jsonTag string, value interface{}) (err error) {
	defer func() {
		if rerr := recover(); rerr != nil {
			err = errors.New(fmt.Sprint(rerr))
		}
	}()
	vOf := reflect.ValueOf(value)
	if vOf.Kind() == reflect.Ptr {
		vOf = vOf.Elem()
	}
	for i := 0; i < valueOf.Len(); i++ {
		item := valueOf.Index(i)
		if item.Kind() == reflect.Ptr {
			item = item.Elem()
		}
		if item.Kind() != reflect.Struct {
			continue
		}

		if err = setStructValue(item, jsonTag, value); err != nil {
			return err
		}
	}
	return nil
}
func getNewValue(valueField reflect.Value, vOf reflect.Value, jsonTag string) (reflect.Value, error) {
	if valueField.Kind() != vOf.Kind() {
		if vOf.Kind() != reflect.Float64 {
			return reflect.Value{}, errors.New(fmt.Sprintf("value of type %v is not assignable to type josn:%v-%v", vOf.Kind(), jsonTag, valueField.Kind()))
		}
		ivOf := vOf.Interface().(float64)
		switch valueField.Kind() {
		case reflect.Int:
			return reflect.ValueOf(int(ivOf)), nil
		case reflect.Int32:
			return reflect.ValueOf(int32(ivOf)), nil
		case reflect.Int64:
			return reflect.ValueOf(int64(ivOf)), nil
		case reflect.Float32:
			return reflect.ValueOf(float32(ivOf)), nil
		default:
			return reflect.Value{}, errors.New(fmt.Sprintf("value of type %v is not assignable to type json:%v ->%v", vOf.Kind(), jsonTag, valueField.Kind()))
		}
	}
	return vOf, nil
}

func getJsonTag(structField reflect.StructField) string {
	fieldName := structField.Name
	tag := structField.Tag.Get("json")

	if tag == "" {
		return fieldName
	}
	if strings.Contains(tag, ",") {
		split := strings.Split(tag, ",")
		if len(split) > 0 {
			s := split[0]
			if s == "" {
				return fieldName
			}
			return s
		}
	}
	return tag
}
