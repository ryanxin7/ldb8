package main

import "fmt"

func main() {
	var a string
	//pair <statictype:string , value: "好大">
	a = "好大"

	//pair <type: string , value: "aceld">
	var alltype interface{}
	alltype = a

	str, _ := alltype.(string)
	fmt.Println(str)
}
