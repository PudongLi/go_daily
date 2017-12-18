package main
/*
递归调用
 */
import "fmt"

func main(){
	testFib()
}

func calc(n int) int {
	if n == 1 {
		return 1
	}
	return calc(n-1)*n
}
func Fibonacci(n int) int{
	//斐波那契数列
	if n <= 1{
		return 1
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
func testFib(){
	for i:=0;i<10;i++{
		n:=Fibonacci(i)
		fmt.Println(n)
	}
}

