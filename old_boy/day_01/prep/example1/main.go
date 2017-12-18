package main

import (
	"fmt"
	"time"
)

const (
	Man = 1
	Female = 2
	Hello = 1

)

func main()  {

	now := time.Now().Unix() //Unix 为1970年到现在的时间戳
	fmt.Println(now)
	if now % 2 == 0 {
		fmt.Println(Female)
	}else {
		fmt.Println(Man)
	}
	example()

}

