package main
/*
编写程序，在终端输出九九乘法表
 */

import (
	"fmt"
)

func main() {
	printMultiplicationTable()
}

func printMultiplicationTable(){

	for i := 1;i < 10;i ++ {

		for j := 1;j <= i;j ++ {
			fmt.Printf("%d * %d = %d\t",j, i, i*j)
		}
		fmt.Println()
	}

}