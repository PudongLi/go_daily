package main

import "fmt"


func add(result int){


	for num1 := 0; num1 <= result; num1 ++{
		fmt.Printf("%d + %d = %d \n", num1, result - num1, result)
	}
}
func main(){
	add(10)
}
