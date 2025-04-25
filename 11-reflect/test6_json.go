package main

// 结构体标签在json中的应用

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title string   `json:"title"`
	Year  int      `json:"year"`
	Price int      `json:"price"`
	Actor []string `json:"actor"`
}

func main() {
	movie1 := Movie{"喜剧之王", 2000, 10, []string{"zhouxingchi", "zhangbaizhi"}}

	//编码过程 结构体-> json

	jsonStr, err := json.Marshal(movie1)
	if err != nil {
		fmt.Println("json marshal error ", err)
		return
	}
	fmt.Printf("jsonStr = %s\n", jsonStr)

	//解码的过程 jsonStr ——> 结构体
	//jsonStr = {"title":"喜剧之王","year":2000,"price":10,"actor":["zhouxingchi","zhangbaizhi"]}
	myMovie := Movie{}
	err = json.Unmarshal(jsonStr, &myMovie)
	if err != nil {
		fmt.Println("json unmarshal error ", err)
		return
	}

	fmt.Printf("%v\n", myMovie)
}

/*
jsonStr = {"title":"喜剧之王","year":2000,"price":10,"actor":["zhouxingchi","zhangbaizhi"]}
{喜剧之王 2000 10 [zhouxingchi zhangbaizhi]}
*/
