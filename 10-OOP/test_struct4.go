package main

import "fmt"

//本质是一个指针

type AnimalIF interface {
	Sleep()
	GetColor() string //获取动物的颜色
	GetType() string  //获取动物的种类
}

//具体的类

type Cat struct {
	color string //猫的颜色
}

func (this *Cat) Sleep() {
	fmt.Println("Cat Sleep")
}

func (this *Cat) GetColor() string {
	return this.color
}

func (this *Cat) GetType() string {
	return "Cat"
}

type Dog struct {
	color string
}

func (this *Dog) Sleep() {
	fmt.Println("Dog Sleep")
}

func (this *Dog) GetColor() string {
	return this.color
}

func (this *Dog) GetType() string {
	return "Dog"
}

func showAnimal(animal AnimalIF) {
	animal.Sleep() //多态
	fmt.Println("color = ", animal.GetColor())
	fmt.Println("type = ", animal.GetType())
}

func main() {
	/*
		var animal AnimalIF //接口的数据类型，父类指针
		animal = &Cat{"Green"}
		animal.Sleep() //调用的就是Cat的Sleep()方法

		animal = &Dog{"Blue"}
		animal.Sleep() //调用Dog的Sleep()方法,多态现象
	*/

	cat := Cat{color: "red"}
	dog := Dog{color: "Yellow"}

	showAnimal(&cat)
	showAnimal(&dog)

	/*
		输出：
		Cat Sleep
		color =  red
		type =  Cat
		Dog Sleep
		color =  Yellow
		type =  Dog
	*/
}
