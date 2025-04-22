package main

import "fmt"

func printArry(myArray []int) {
	//引用传递
	//_ 表示匿名的变量

	for _, value := range myArray {
		fmt.Println("value =", value)
	}

	myArray[0] = 100
}
func main() {
	myArray := []int{1, 2, 3, 4, 5} //动态数组，切片slice

	fmt.Printf("myArray type= %T\n", myArray)
	printArry(myArray)

	fmt.Println("====")

	for _, value := range myArray {
		fmt.Println("value =", value)
	}

}
