package main

//定义结构体标签
import (
	"fmt"
	"reflect"
)

type resume struct {
	Name string `info:"name" doc:"我的名字"`
	Sex  string `info:"sex"`
}

func findTag(str interface{}) {
	t := reflect.TypeOf(str).Elem()
	for i := 0; i < t.NumField(); i++ {
		taginfo := t.Field(i).Tag.Get("info")
		tagdoc := t.Field(i).Tag.Get("doc")
		fmt.Printf("info：%s doc: %s\n", taginfo, tagdoc)
	}
}

func main() {
	var re resume
	findTag(&re)
}

/*
info：name doc: 我的名字
info：sex doc:
*/
