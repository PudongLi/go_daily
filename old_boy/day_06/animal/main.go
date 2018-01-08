package main

import (
	"fmt"
)

type Animal interface {
	Eat()
	Talk()
}

type Cat struct {
	name string

}

func (c *Cat)Eat()  {
	fmt.Println(c.name, "cat eating")
}

func (c *Cat)Talk(){
	fmt.Println(c.name, "cat talking")
}

type Dog struct {
	name string

}

func (d *Dog)Eat()  {
	fmt.Println(d.name, "dog eating")
}

func (d *Dog)Talk(){
	fmt.Println(d.name, "dog talking")
}

func  TestOperator() {
	var animalList []Animal
	d := &Dog{
		name:"wawa",
	}
	animalList=append(animalList, d)
	d1 := &Dog{
		name:"heihei",
	}
	animalList = append(animalList, d1)
	c1 := &Cat{
		name:"xiaohua",
	}
	animalList = append(animalList, c1)

	for _,v := range animalList{
		v.Eat()
		v.Talk()
	}

	}

func test(){
	var a Animal
	var d Dog
	d.Eat()
	a = &d
	a.Eat()
}
func main() {
	TestOperator()
}


