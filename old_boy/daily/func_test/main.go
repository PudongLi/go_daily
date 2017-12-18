package main

import "fmt"

func main() {
	Print(1,2,"hello","world")

}


func Print(args ...interface{}){
	for _, arg := range args{
		fmt.Println(arg)
	}
}