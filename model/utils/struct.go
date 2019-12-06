package commutils

import (
	"encoding/json"
	"gopkg.in/jeevatkm/go-model.v1"
	"reflect"
	"strings"
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

// 结构体转换成Map;
func StructToMap(obj interface{}) map[string]interface{} {
	var data = make(map[string]interface{})
	s := reflect.ValueOf(obj).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		if strings.HasPrefix(typeOfT.Field(i).Name, "XXX_") {
			continue
		}
		f := s.Field(i)
		data[typeOfT.Field(i).Name] = f.Interface()
	}
	return data
}

// map转换成结构体
func MapToStruct(input interface{}, obj interface{}) (err error) {
	// 序列化;
	j, errM := json.Marshal(input)
	if errM != nil {
		err = errM
	}
	// 反序列化;
	return json.Unmarshal([]byte(j), obj)
}
