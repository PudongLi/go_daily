package main
/*
输出101~200之间的素数并统计
 */
 import (
 	"fmt"
 )

func main() {
	var prime [] int
	for i := 101;i < 200;i++ {
		n := 0
		for j := 2;j < i;j++ {
			if i % j == 0 {
				n ++
			}
		}
		if n < 1 {
			prime = append(prime, i)
		}
	}
	fmt.Printf("101~200之间有%d个素数\n", len(prime))
	fmt.Printf("它们是:%v\n", prime)
}
