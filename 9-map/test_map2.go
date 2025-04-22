package main

import "fmt"

func printMap(cityMap map[string]string) {
	//cityMap是一个引用传递
	for key, value := range cityMap {
		fmt.Println("key= ", key)
		fmt.Println("value= ", value)

	}
}

func ChangeValue(cityMap map[string]string) {
	cityMap["China"] = "chongqing"
}

func main() {
	cityMap := make(map[string]string)

	//添加
	cityMap["China"] = "Beijing"
	cityMap["Japan"] = "Tokyo"
	cityMap["USA"] = "New York"

	//遍历
	for key, value := range cityMap {
		fmt.Println("key= ", key)
		fmt.Println("value= ", value)
	}

	//删除
	delete(cityMap, "USA")

	//修改
	cityMap["China"] = "xian"

	fmt.Println("---------------")
	ChangeValue(cityMap)
	//遍历
	printMap(cityMap)

}
