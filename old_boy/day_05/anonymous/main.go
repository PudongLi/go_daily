package main

import (
	"fmt"
)
/*
匿名变量，实现继承
 */
type People struct {
	Name string
	Age int
	}

type Student struct {
	Score int
	People
}
func (s *Student)Set(name string, age int){
	s.Name = name
	s.Age = age
	return

}
func main() {
	var s Student
	s.Name = "lipd"
	s.Age = 23
	s.Score = 100

	fmt.Printf("%#v\n", s)

	var k Student
	k.Set("wyr", 23)
	fmt.Printf("%#v\n", k)
}
