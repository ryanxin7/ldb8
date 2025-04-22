package main

import "fmt"

func main() {

	//第一种方式
	//声明myMao1 是一种map类型  key是 string  ， value 是 string
	var myMap1 map[string]string
	if myMap1 == nil {
		fmt.Println("myMap1 is 空 map")
	}

	//在使用mao前，需要先用make给map分配数据空间
	myMap1 = make(map[string]string, 10)

	myMap1["one"] = "java"
	myMap1["two"] = "go"
	myMap1["three"] = "python"

	fmt.Println(myMap1)

	//第二种方式,声明

	myMap2 := make(map[int]string)
	myMap2[1] = "java"
	myMap2[2] = "go"
	myMap2[3] = "python"

	fmt.Println(myMap2)

	//第三种方式,声明
	//初始化
	myMap3 := map[string]string{
		"one":   "java",
		"two":   "go",
		"three": "python",
	}

	fmt.Println(myMap3)

}
