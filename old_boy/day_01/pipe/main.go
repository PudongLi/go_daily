package main
import (
	"fmt"

)

func main(){
	pipe := make(chan int, 3) //声明一个管道，管道内元素的类型为int，3为管道内元素的个数
	pipe <- 1
	pipe <- 2
	a := <- pipe

	b := <- pipe
	fmt.Printf("%d\n", a)
	fmt.Printf("%d\n", b)
	Hello()
}

