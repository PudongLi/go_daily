package main

import (
	"fmt"
	"encoding/json"
)

type Book struct {
	Title       string
	Author      []string
	Publisher   string
	Price       float64
	IsPublished bool
}

func main(){
	b := []byte (`{
	"Title":"go programming language",
    "Author":["john", "ada", "aclice"],
    "Publisher":"qinghua",
    "IsPublished":true,
    "Price":99
	}`)

	var book Book
	err := json.Unmarshal(b, &book)
	if err !=nil{
		fmt.Println("error in translating,", err.Error())
		return
	}
	fmt.Println(book.Author)



	}

