package funct

import (
	"encoding/json"
	"reflect"
	"strconv"
)

//	type test struct {
//	    Name     string `json:"name" default:"bbc"`
//	    Addr     string `json:"addr"`
//	    Port     uint   `json:"port" default:"88"`
//	    User     string `json:"user"`
//	    Password string `json:"password"`
//	}
//
// data, err := MarshalJSON(t)
//
//	if err != nil {
//	    panic(err)
//	}
func MarshalJSON(i interface{}) ([]byte, error) {
	typeof := reflect.TypeOf(i)
	valueof := reflect.ValueOf(i)
	for i := 0; i < typeof.Elem().NumField(); i++ {
		if valueof.Elem().Field(i).IsZero() {
			def := typeof.Elem().Field(i).Tag.Get("default")
			if def != "" {
				switch typeof.Elem().Field(i).Type.String() {
				case "int":
					result, _ := strconv.Atoi(def)
					valueof.Elem().Field(i).SetInt(int64(result))
				case "uint":
					result, _ := strconv.ParseUint(def, 10, 64)
					valueof.Elem().Field(i).SetUint(result)
				case "string":
					valueof.Elem().Field(i).SetString(def)
				}
			}
		}
	}
	return json.Marshal(i)
}
