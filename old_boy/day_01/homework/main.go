/*
day01作业：打印输出字符串、二进制、十进制、十六进制、浮点型
 */
package main

import (
	"fmt"
)

func main(){
	var varStr string
	varStr = "15"

	varBinary := 15

	varDecimal := 15

	varHexadecimal := 15

	var varFloat float32
	varFloat = 15.00

	fmt.Printf("%s\n", varStr)
	fmt.Printf("%b\n", varBinary)
	fmt.Printf("%d\n", varDecimal)
	fmt.Printf("%x\n", varHexadecimal)
	fmt.Printf("%g\n", varFloat)



}

