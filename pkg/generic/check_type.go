package generic

import (
	errors2 "github.com/zepeng-jiang/go-basic-demo/pkg/errors"
	"reflect"
)

// CheckType ,check input value type.
// Resolve Go not have Generic.
func CheckType(typeOf string, val interface{}) error {
	t := reflect.TypeOf(val).Kind().String()
	if t == "struct" || t == "ptr" {
		sName := reflect.TypeOf(val).String()
		if sName != typeOf {
			return errors2.InvalidTypeError
		}
		return nil
	}
	if typeOf != t {
		return errors2.InvalidTypeError
	}
	return nil
}
