package main

import (
	"os"
	"fmt"
)

func main() {
	hostName ,err := os.Hostname()
	if err != nil{
		fmt.Printf("%v\n", err)
	}
	env := os.Getenv("gopath")
	fmt.Printf("hostName %s env %s\n", hostName, env)
	a := 3
	b := 4
	a, b = swap(a, b)
	fmt.Printf("%d %d",a, b)
}
func swap(a int, b int)(int, int){
	a,b = b,a
	return a, b

}
