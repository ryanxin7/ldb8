package main

import "fmt"

func swap(a, b *int) {
	var temp int
	temp = *a
	*a = *b
	*b = temp

}
func main() {
	var a, b int = 10, 20
	swap(&a, &b)

	fmt.Println("a=", a, "b=", b)

	//二级指针

	var p *int
	p = &a
	fmt.Println(&p)
	fmt.Println(p)

	var pp **int
	pp = &p
	fmt.Println(&p)
	fmt.Println(pp)
}
