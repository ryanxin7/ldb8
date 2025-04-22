package main

//import 导包

import (
	_ "ldb8/5-init/lib1" //给fmt包起一个别名，匿名，无法使用当前包的方法，但是会执行当前的包内部的`init()`方法
	//mylib2 "ldb8/5-init/lib2"  //别名
	. "ldb8/5-init/lib2" //直接导入当前包里
)

func main() {
	//lib1.Lib1Test()
	//mylib2.Lib2Test()

	Lib2Test()

}
