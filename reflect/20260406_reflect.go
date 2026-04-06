package main

import (
	"fmt"
	"reflect"
)

// 使用场景：
// 1.万能的 JSON 序列化 (json.Marshal)
// 2.数据库 ORM (映射)
// 3.配置文件加载 (Viper / Env)

//补充：
// 1.首选泛型，次选断言，底选反射
// 2.泛型：用于编写适用于“多个类型”的“同一处理逻辑”的通用方法
// 3.断言：用于编写适用于“多个类型”的“不同处理逻辑”的通用方法
// 4.反射：用于编写适用于“所有类型”的“同一处理逻辑”的通用方法

type User struct {
	ID   int    `json:"id" model:"pk"`
	Name string `json:"name"`
}

// 读
func readStruct(s interface{}) {
	v := reflect.ValueOf(s)
	t := reflect.TypeOf(s)

	// 确保传入的是结构体
	if t.Kind() != reflect.Struct {
		return
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)             // 类型信息（名字、Tag等）
		value := v.Field(i).Interface() // 实际数值

		tag := field.Tag.Get("json") // 获取 Json Tag
		fmt.Printf("字段: %s, Tag: %s, 值: %v\n", field.Name, tag, value)
	}
}

// 写
func modifyValue(s interface{}) {
	v := reflect.ValueOf(s)

	// 1. 检查是否为指针
	if v.Kind() != reflect.Ptr {
		return
	}

	// 2. 获取指针指向的元素
	e := v.Elem()

	// 3. 只有导出的字段（大写）才能被修改
	nameField := e.FieldByName("Name")
	if nameField.CanSet() {
		nameField.SetString("New Name")
	}
}

func main() {
	u := User{
		ID:   123,
		Name: "WTF",
	}
	readStruct(u)
	modifyValue(&u)
	fmt.Println(u)
}
