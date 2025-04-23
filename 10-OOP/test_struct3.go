package main

import "fmt"

type Human struct {
	name string
	sex  string
}

func (this *Human) Eat() {
	fmt.Println("Human.eat()...")
}

func (this *Human) Walk() {
	fmt.Println("Human.Walk()...")
}

type SuperMan struct {
	Human //SuperMan类继承了Human类的方法
	level int
}

// 重定义父类的方法Eat()
func (this *SuperMan) Eat() {
	fmt.Println("SuperMan.Eat()...")
}

// 子类的新方法
func (this *SuperMan) Fly() {
	fmt.Println("SuperMan.Fly()...")
}

func (this *SuperMan) Show() {
	fmt.Println("name =", this.name)
	fmt.Println("sex =", this.sex)
	fmt.Println("level =", this.level)
}

func main() {
	h := Human{name: "zhang3", sex: "female"}
	h.Eat()
	h.Walk()

	//定义一个子类对象
	//s := SuperMan{Human{"zhao6", "male"}, 99}

	var s SuperMan
	s.name = "zhao6"
	s.sex = "male"
	s.level = 99

	s.Walk() //父类的方法
	s.Eat()  //子类的方法
	s.Fly()  // 子类的方法
	s.Show()
}
