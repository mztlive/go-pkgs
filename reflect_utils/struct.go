package reflect_utils

import (
	"fmt"
	"reflect"
	"time"

	"github.com/mztlive/go-pkgs/string_utils"
)

// GetNameFromStruct 获取结构体的名字
func GetNameFromStruct(s any) string {
	// 通过反射获取结构体的名字
	val := reflect.ValueOf(s)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	return val.Type().Name()
}

// GetSnakeNameFromStruct 获取结构体的名字(蛇形命名)
func GetSnakeNameFromStruct(s any) string {
	// 通过反射获取结构体的名字
	val := reflect.ValueOf(s)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	return string_utils.ToSnakeCase(val.Type().Name())
}

// StructToMapStr converts a struct to map[string]string.
// The map keys are taken from struct field tags "db" and "json".
//
// Rules:
// - Fields with empty tag will be skipped.
// - For pointer fields, nil value will be skipped.
// - All field values will be converted to string.
// - time.Time will be converted to "2006-01-02 15:04:05" format string.
// - bool/pointer to bool will be converted to "0" or "1".
//
// Example:
//
//	type User struct {
//	  Name string `json:"name"`
//	  Age  int    `json:"age"`
//	}
//
//	user := User{"John", 30}
//	m := StructToMapStr(user)
//	// m = map[string]string{"name":"John", "age":"30"}
func StructToMapStr(s any) map[string]string {
	data := make(map[string]string)
	structToMapStr(reflect.ValueOf(s), data)
	return data
}

func structToMapStr(v reflect.Value, data map[string]string) {
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		if value.Kind() == reflect.Ptr && value.IsNil() {
			continue
		}

		if value.Kind() == reflect.Ptr {
			value = value.Elem()
		}

		if value.Kind() == reflect.Struct && value.Type() != reflect.TypeOf(time.Time{}) {
			structToMapStr(value, data)
			continue
		}

		tag := field.Tag.Get("db")
		if tag == "" {
			tag = field.Tag.Get("json")
		}

		if tag != "" {
			strValue := fmt.Sprint(value.Interface())
			if value.Type() == reflect.TypeOf(time.Time{}) || value.Type() == reflect.TypeOf(&time.Time{}) {
				strValue = value.Interface().(time.Time).Format("2006-01-02 15:04:05")
			} else if value.Kind() == reflect.Bool {
				if value.Bool() {
					strValue = "1"
				} else {
					strValue = "0"
				}
			}

			data[tag] = strValue
		}
	}
}
