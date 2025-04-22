package main

import "fmt"

func main() {
	var numbers = make([]int, 3, 5)
	fmt.Printf("len = %d,cap = %d, slice = %v", len(numbers), cap(numbers), numbers)

	// 向numbers切片追加一个元素1，number len = 4 ，[0,0,0,1] , cap = 5

	fmt.Println()
	numbers = append(numbers, 1)

	fmt.Printf("len = %d,cap = %d, slice = %v", len(numbers), cap(numbers), numbers)

	fmt.Println()
	numbers = append(numbers, 2)

	fmt.Printf("len = %d,cap = %d, slice = %v", len(numbers), cap(numbers), numbers)

	fmt.Println()

	// 向一个容量cap已经满了的slice追加元素
	numbers = append(numbers, 3)

	fmt.Printf("len = %d,cap = %d, slice = %v", len(numbers), cap(numbers), numbers)

	fmt.Println()

	var numbers2 = make([]int, 3)
	fmt.Printf("len = %d,cap = %d, slice = %v", len(numbers2), cap(numbers2), numbers2)
	fmt.Println()
	numbers2 = append(numbers2, 1)
	fmt.Printf("len = %d,cap = %d, slice = %v", len(numbers2), cap(numbers2), numbers2)
	fmt.Println()
	var numbers3 = make([]int, 4)
	fmt.Printf("len = %d,cap = %d, slice = %v", len(numbers3), cap(numbers3), numbers3)

	fmt.Println()
	numbers3 = append(numbers3, 5)
	fmt.Printf("len = %d,cap = %d, slice = %v", len(numbers3), cap(numbers3), numbers3)

}
