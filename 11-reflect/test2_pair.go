package main

import "fmt"

type Reader interface {
	ReadBook()
}

type Writer interface {
	WriteBook()
}

// 具体类型
type Book struct {
}

func (b Book) ReadBook() {
	fmt.Println("Read Book ...")
}

func (b Book) WriteBook() {
	fmt.Println("Write Book ...")
}

func main() {
	// b : pair <type:Book, value:book{} 地址>
	b := &Book{}

	//r : pair<type:,value:>
	var r Reader
	// r: pair<type:Book,value:book{}地址>

	r = b
	r.ReadBook()

	var w Writer
	//w: pair<type:Book,value:book{}地址>

	w = r.(Writer)

	w.WriteBook()
}
