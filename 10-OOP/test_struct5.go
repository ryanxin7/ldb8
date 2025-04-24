package main

import (
	"fmt"
)

//interface{}是万能数据类型

func myFunc(arg interface{}) {
	fmt.Println("myFunc is called ...")
	fmt.Println(arg)

	value, ok := arg.(string)
	if !ok {
		fmt.Println("arg is not string type")
	} else {
		fmt.Println("arg is string type, value = ", value)
		fmt.Printf("value type is %T\n", value)
	}

}

//interface {} 如何区分此时引用的底层数据类型到底是什么？

// 给 interface{} 通过 ”类型断言“ 的机制

type Books struct {
	auth string
}

func main() {
	books := Books{auth: "Golang"}

	myFunc(books)
	myFunc(100)
	myFunc(false)
	myFunc(3.12)
	myFunc("好大")
}
