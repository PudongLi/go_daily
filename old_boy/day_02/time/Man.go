package main

import (
	"time"
	"fmt"
)

const (
	Man = 1
	Female = 2
)

func main() {
	t := time.Now().Unix()
	if  t % Female == 0{
		fmt.Println("female")
	}else {
		fmt.Println("man")
	}
}
