package main

import "fmt"

// 声明一种新的数据类型 myint  ，是 int 的一个别名
type Myint int

var x Myint = 10

// 定义一个结构体
type Book struct {
	title  string
	author string
}

func changeBook(book Book) {
	//传递一个book的副本
	book.author = "li4"
}

func changeBook2(book *Book) {
	//传递一个book的副本
	book.author = "li4"
}

func main() {

	fmt.Println(x)
	/*
		var a myint = 10
		fmt.Println("a=", a)
		fmt.Printf("type of a =%T\n", a)

		a= 10
		type of a =main.myint
	*/

	var book1 Book
	book1.title = "Golang"
	book1.author = "zhang3"

	fmt.Printf("%v\n", book1)

	changeBook(book1)
	fmt.Printf("%v\n", book1)
	changeBook2(&book1)
	fmt.Printf("%v\n", book1)

	/*
		输出：
		{Golang zhang3}
		{Golang zhang3}
		{Golang li4}
	*/
}
