package main

import "fmt"

func main() {
	s := []int{1, 2, 3} //len= 3 ,cap = 3 ,[1,2,3]

	//[0,2)
	s1 := s[0:2]
	fmt.Println(s1)

	s2 := s[1:2]
	fmt.Println(s2)

	s3 := s[1]
	fmt.Println(s3)

	s1[0] = 100
	fmt.Println(s1)

	//copy 可以将底层数组的slice 一起拷贝

	s4 := make([]int, 3) // s2 = [0,0,0]

	//将s中的值，依次拷贝到s4中
	copy(s4, s)

	fmt.Println(s4)
}
