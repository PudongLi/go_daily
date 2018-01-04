package main
/*
结构体

 */
import "fmt"

type Student struct{
	Age int
	Name string
	Sex string
	Grade string
	Score int
}

func testStruct(){
	var s Student
	s.Age = 23
	s.Name = "lipd"
	s.Sex = "男"
	s.Grade = "未知"
	s.Score = 100
	fmt.Println(s)
}

func main() {
	testStruct()
}