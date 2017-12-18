package main

import "fmt"
import "time"

func PrintOdd(n int){
	for i := 0; i < n; i++ {
		if i % 2 == 0{
			fmt.Printf("odd %d\n", i)
		}
	}
}

func PrintEven (n int)  {
	for i := 0; i < n; i++ {
		if i % 2 != 0{
			fmt.Printf("even %d\n", i)
		}
	}
}
func main(){
	//使用go调用函数后，调用的函数将是独立调度的，主线程执行完毕后，printodd还为执行完成
	go PrintOdd(10)
	go PrintEven(10)
	time.Sleep(1*time.Second)
	//PrintEven(10)
	//PrintOdd(10)
}
