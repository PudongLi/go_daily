package main
/*
链表，前插法 后插法
通过next 访问下一个节点
构造函数
 */
import (
	"fmt"
	)

type Student struct {
	Age int
	Name string
	Next *Student
}
/*
手动实现构造函数
 */
func newStudent(age int, name string) *Student{
	return  &Student{
		Age:age,
		Name:name,
	}
}

func testList(){
	var s Student
	s.Age = 20
	s.Name = "a"
	p := new(Student)
	p.Age = 21
	p.Name = "b"
	s.Next = p
	}

func insertListHeader(header *Student,age int, name string) *Student{
	var a = &Student{}
	a.Name = name
	a.Age = age
	a.Next = header
	return a
}

func insertListEnd(){

}

func main() {
	//testList()
	var header=&Student{}
	header = insertListHeader(header, 10, "a")
	header = insertListHeader(header, 11, "b")
	header = insertListHeader(header, 12, "c")
	fmt.Println(*header)

}