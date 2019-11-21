package commutils

import (
	"gopkg.in/jeevatkm/go-model.v1"
	"reflect"
)

func IsZero(v interface{}) bool {
	return model.IsZero(v)
}

func IsZeroOrNil(v interface{}) bool {
	if v == nil {
		return true
	}
	if reflect.ValueOf(v).IsNil() {
		return true
	}
	return model.IsZero(v)
}
