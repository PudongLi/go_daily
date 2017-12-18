package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "hello world"
	index1 := strings.Index(str1, "l")
	fmt.Printf("%v\n", index1)

}