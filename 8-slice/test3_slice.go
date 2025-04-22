package main

import "fmt"

func main() {
	//1.声明slice1 是一个切片，并且初始化，默认值1，2，3，4，5。长度len 是5
	//slice1 := []int{1, 2, 3, 4, 5}

	//2.声明slice1 是一个切片，但是没有给slice分配空间
	var slice1 []int
	//slice1 = make([]int, 5) //开辟5个空间，默认值0

	//3.声明slice1 是一个切片，同时给slice 分配空间，5个空间初始值是0
	//var slice1 []int = make([]int, 5)

	//4.声明slice1 是一个切片，同时给slice 分配空间，5个空间初始值是0,通过:= 推导出slice 是一个切片
	//slice1 := make ([]int,5)

	//slice1[0] = 100

	//5.判断一个slice是否是空切片
	if slice1 == nil {
		fmt.Println("slice1 是一个空切片")
	} else {
		fmt.Println("slice1 是有空间的")
	}

	fmt.Printf("len = %d, slice = %v\n", len(slice1), slice1)

}
