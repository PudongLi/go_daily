package main

import "fmt"
var BorrowBookMap map [string]([]string)
func main() {

	BorrowBookMap = make(map [string]([]string))

	flag :=AddBook("他改变了中国", 10,"xixix", "2017/01/01")
	if flag{
		fmt.Println("插入成功")
	}
	fmt.Println(*(Books[0]))
	AddBook("他改变了世界", 100,"xixix", "2017/01/01")
	AddBook("java从入门到放弃", 22,"no one", "2018/01/01")
	fmt.Println("-----------------------")
	SelectBookByTitle("java从入门到放弃")
	SelectBookByTitle("大腕")
	SelectBookByAuthor("xixix")

	stu1 := NewStudent("lipd","123", "nan", "9",[]string{})

	stu2 := NewStudent("mark","456", "nan", "9",[]string{})
	PrintStudent(stu1)
	stu1.BorrowBook("他改变了中国")
	stu1.BorrowBook("他改变了世界")
	stu2.BorrowBook("他改变了世界")

	PrintStudent(stu1)
	PrintAllBook()
	fmt.Println("-----------------------")
	PrintBorrowRecord()
	fmt.Println("-----------------------")
	PrintAllBook()




}