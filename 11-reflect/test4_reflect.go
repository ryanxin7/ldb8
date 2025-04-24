package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (this User) Call() {
	fmt.Println("user is call...")
	fmt.Println("%v\n", this)
}

func main() {
	user := User{1, "张飞", 18}

	DofiledAndMethod(user)
}

func DofiledAndMethod(input interface{}) {
	//获取 input 的 type
	inputType := reflect.TypeOf(input)
	fmt.Println("inputType is :", inputType.Name())

	// 获取input的value

	inputValue := reflect.ValueOf(input)
	fmt.Println("inputValue is :", inputValue)

	//通过type 获取里面的字段
	//1.获取interface的reflect.Type ,通过Type 得到Numfield ,进行便利
	//2.得到每个field ，数据类型
	//3.通过filed有一个Interface（）方法得到对应的value

	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		value := inputValue.Field(i).Interface()

		fmt.Printf("%s : %v = %v\n", field.Name, field.Type, value)
	}

	//通过type 获取里面的方法。调用
	for i := 0; i < inputType.NumMethod(); i++ {
		m := inputType.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}

}
