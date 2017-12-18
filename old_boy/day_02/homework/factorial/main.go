package main
/*
计算1~n的阶乘之和
 */
import (
	"fmt"
)
func main(){
	var num uint64 = 3
	var result uint64 = 0
	var i uint64 = 1
	for i <= num {
		result += factorial(i)
		i ++
	}
	fmt.Printf("result %d\n", result)
}

func factorial(n uint64) (result uint64){
	if n > 0{
		result = n * factorial(n-1)
		return result
	}
	return 1

}
