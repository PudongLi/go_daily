package main

import "fmt"

var (
	Result1 int
)

func main(){

	Result1 = funcA()
	fmt.Println(Result1)
}

func funcA() int {
	x := 5
	defer func(){
		x += 1
	}()
	return x

}

