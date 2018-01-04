package main

import "fmt"

/*
书籍
 */
var(
	Books = []*Book{}
)
type Book struct {
	BookTitle string //书名
	Duplication int  //副本数
	Author string    //作者
	PublishDate string   //出版日期

}
/*
打印书籍详情
 */
func PrintBook(book *Book){
	fmt.Println("图书名称：",book.BookTitle," 副本数：",book.Duplication," 作者：",book.Author," 出版日期：",book.PublishDate)

}

/*
打印所有图书详情
 */
func PrintAllBook(){
	for i := 0;i <len(Books);i++{
			PrintBook(Books[i])
	}
}
/*
Book构造函数
 */

func NewBook(title string, duplication int, author string, publishDate string) *Book{
	return &Book{
		BookTitle:title,
		Duplication:duplication,
		Author:author,
		PublishDate:publishDate,
	}
}







