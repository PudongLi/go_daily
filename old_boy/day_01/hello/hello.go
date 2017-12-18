package main

import (
	"fmt"
	"old_boy/day_01/calc"
	mylib "old_boy/day_01/lib"
)
func init(){
	fmt.Println("hello init")
}
func main(){
	Name := "wangyr"
	mylib.SayHello()
	fmt.Println(mylib.Name)
	fmt.Println(Name)
	sum, sub := calc.Calc(10, 9)
	fmt.Printf("%d %d", sum, sub)

}