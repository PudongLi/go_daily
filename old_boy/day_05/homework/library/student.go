package main

import "fmt"

/*
学生
 */


type Student struct {

	StuName string //
	Grade string
	Id string
	Gender string
	Books []string

	}

func PrintStudent(s *Student){
	fmt.Println("学生姓名：",s.StuName," 学生编号：",s.Id," 学生性别：",s.Gender," 年级：",s.Grade," 所借阅的图书",s.Books)

}
func NewStudent(name string, grade string, id string, gender string, books []string) *Student{
	return &Student{
		StuName:name,
		Grade:grade,
		Id:id,
		Gender:gender,
		Books:books,
		}
}
