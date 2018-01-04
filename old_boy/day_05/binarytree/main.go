package main

import (
	"fmt"

)

type Student struct {
	Name string
	Age int
	Score float32
	left *Student
	right *Student
	}
/*
前序遍历
 */
func trans(root *Student){
	if root==nil{
		return
	}
	fmt.Println(root)
	trans(root.left)
	trans(root.right)
}

func main(){
	var root *Student = new(Student)
	root.Name = "tony"
	root.Age = 19
	root.Score = 100
	root.left = nil
	root.right = nil

	var left1 *Student = new(Student)
	left1.Name = "stu2"
	left1.Age = 19
	left1.Score = 99
	root.left = left1

	var right1 *Student = new(Student)
	right1.Name = "stu3"
	right1.Age = 20
	right1.Score = 98
	root.right = right1

	trans(root)


	}

