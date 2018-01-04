package main

import "fmt"
/*
添加图书
 */
func AddBook(title string, duplication int, author string, pubulishDate string) bool{
	myBook := NewBook(title, duplication, author, pubulishDate)
	//先判断书籍是否在列表内
	lenBooks := len(Books)
	if duplication <= 0{
		fmt.Println("新增书籍的副本数不能小于0")
		return false
	}
	if lenBooks > 0{
		for i := 0;i <lenBooks;i++{
			if (*(Books[i])).BookTitle == myBook.BookTitle{
				fmt.Println("此书已存在，请勿重复添加")
				return false
			}
		}
	}
	Books = append(Books, myBook)
	return true
}

/*
根据书籍名称查询图书
 */
func SelectBookByTitle(title string) bool{
	for i := 0;i <len(Books);i++{
		if (*(Books[i])).BookTitle == title{
			PrintBook(Books[i])
			return true
		}
	}
	fmt.Println("对不起，没有您要查找的书籍")
	return false
}

/*
根据书籍作者查询图书
 */

func SelectBookByAuthor(author string) bool{
	flag := false
	for i := 0;i <len(Books);i++{
		if (*(Books[i])).Author == author{
			PrintBook(Books[i])
			flag = true
		}
	}
	if !flag  {
		fmt.Println("对不起，没有您要查找的书籍")
	}
	return flag

}
/*
根据书籍出版日期查询图书
 */

func SelectBookByPublishDate(publishDate string) bool{
	flag := false
	for i := 0;i <len(Books);i++{
		if (*(Books[i])).PublishDate == publishDate{
			PrintBook(Books[i])
			flag = true
		}
	}
	if !flag  {
		fmt.Println("对不起，没有您要查找的书籍")
	}
	return flag

}

/*
借阅图书
 */
func (s *Student) BorrowBook(bookName string) bool{
	flag := SelectBookByTitle(bookName)
	if flag {
		for i := 0;i <len(Books);i++{
			if (*(Books[i])).BookTitle == bookName{
				if Books[i].Duplication <= 0{
					fmt.Println("此书籍库存不足，不可借阅")
					return false
				}
				Books[i].Duplication = Books[i].Duplication - 1
			}
		}
		s.Books = append(s.Books, bookName)
		RecortBorrow(s.StuName, bookName)
	}
	return true
}

/*
记录图书借阅
 */
func RecortBorrow(stuName string, bookName string){
	stuNames, exist := BorrowBookMap[bookName]
	if exist{
		stuNames = append(stuNames, stuName)
		BorrowBookMap[bookName] = stuNames
	}else {
		BorrowBookMap[bookName] = []string{stuName}
	}

}

func PrintBorrowRecord(){
	for bookName, stuNames := range BorrowBookMap{
		fmt.Println("图书名称：",bookName," 借阅人：", stuNames)
	}
}