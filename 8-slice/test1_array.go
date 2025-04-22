package main

import "fmt"

func printArray(myArray [4]int) {
	for index, value := range myArray {
		fmt.Println("index=", index, "value=", value)
	}
}
func main() {
	//固定长度的数组
	var myArray [10]int

	myArray2 := [10]int{1, 2, 3, 4}
	myArray4 := [4]int{1, 2, 3, 4}

	//for i := 0; i < 10; i++ {
	for i := 0; i < len(myArray); i++ {
		fmt.Println(myArray[i])
	}

	for index, value := range myArray2 {
		fmt.Println("index = ", index, "value=", value)
	}

	//查看数组的数据类型

	fmt.Printf("myArray1 types= %T\n", myArray)
	fmt.Printf("myArray2 types= %T\n", myArray2)
	fmt.Printf("myArray4 types= %T\n", myArray4)

	printArray(myArray4)

}
