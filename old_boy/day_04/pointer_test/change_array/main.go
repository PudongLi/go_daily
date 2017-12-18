package main

import "fmt"

func main()  {
	a := [6]int{1,2,3,4,5,6}
	changeArray(&a)
	fmt.Println(a)
}

func changeArray(a *[6]int){
	(*a)[0] = 1000

}
