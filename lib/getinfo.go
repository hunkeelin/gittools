package kgittools

import (
	"errors"
	"reflect"
	"strings"
)

func GetInfo(s interface{}, ss string) (reflect.Value, error) {
	var typ reflect.Type
	var val reflect.Value
	ptyp := reflect.TypeOf(s)  // a reflect.Type
	pval := reflect.ValueOf(s) // a reflect.Value
	if ptyp.Kind() == reflect.Ptr {
		typ = ptyp.Elem()
		val = pval.Elem()
	} else {
		typ = ptyp
		val = pval
	}
	var r reflect.Value
	for i := 0; i < typ.NumField(); i++ {
		vfld := val.Field(i)
		sfld := typ.Field(i)
		if strings.ToLower(ss) == strings.ToLower(sfld.Name) {
			r = vfld
			return vfld, nil
		}
	}
	return r, errors.New("No fields with " + ss)
}
