package main

import "fmt"

/*
给定一个数字n，求出所有两两相加等于n的组合。
 */

func main() {
	LoopAdd(5)
}

func LoopAdd(num int){
	for i := 0;i <= num;i ++{
		fmt.Printf("%d + %d = %d \n",i, num - i, num)
	}

}