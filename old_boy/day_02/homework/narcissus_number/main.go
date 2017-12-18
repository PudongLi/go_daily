package main
/*
打印100~999之间的水仙花数
 */
import (
	"fmt"
)

func main() {
	for i := 100; i <= 999; i ++ {
		x := i / 100
		y := i % 100 / 10
		z := i % 10
		if (x * x *x + y * y * y + z * z * z) == i{
			fmt.Printf("%d\n", i)
		}
	}
}
