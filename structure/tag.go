package structure

import "reflect"

// GetTagVar 返回结构体中的变量
// 传入标签的名称，通过反射返回对应的变量
func GetTagVar(obj any, varName, tagName string) interface{} {
	// 获取结构体类型
	t := reflect.TypeOf(obj)
	if t.Kind() != reflect.Ptr {
		// 如果传入的参数不是指针类型，则将其转换为指针类型
		objPtr := reflect.New(t)
		objPtr.Elem().Set(reflect.ValueOf(obj))
		t = objPtr.Type()
		obj = objPtr.Interface()
	}

	// 遍历结构体的字段，查找标签的值是否等于传入的 varName
	for i := 0; i < t.Elem().NumField(); i++ {
		field := t.Elem().Field(i)
		if field.Tag.Get(tagName) == varName {
			// 如果找到了对应的字段，就返回该字段的值
			value := reflect.ValueOf(obj).Elem().Field(i)
			return value.Interface()
		}
	}

	return nil
}
