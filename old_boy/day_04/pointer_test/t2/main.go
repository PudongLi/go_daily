package main

import "fmt"
var (
	a = "111"
)

func n(){
	fmt.Println(a)
}

func m(){
	a = "123"
	fmt.Println(a)

}
func main() {
	n()
	m()
	n()

}

