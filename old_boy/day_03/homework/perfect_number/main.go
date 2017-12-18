package main
/*
	给定一个范围，输出这个范围内的所有完数
 */
import (
	"fmt"
)

func main() {
	printPerfectNumber(1, 1000)

}

func printPerfectNumber(startNum int, endNum int ){
	/*
	parameter
		:startNum   起始位置
	    :endNum     结束位置
	 */
	 if startNum <= 0 || endNum <= 0{
	 	fmt.Println("参数不能小于等于0")
	 	return
	 }
	 if startNum >= endNum {
	 	fmt.Println("参数2要大于参数1")
	 	return
	 }

	for i := startNum; i < endNum; i++ {
		currentSum := 0
		for j := 1; j <= i/2; j++ {
			if i%j == 0 {
				currentSum += j
			}
		}
		if currentSum == i {
			fmt.Printf("%d is factory number\n", i)
		}
	}
}


