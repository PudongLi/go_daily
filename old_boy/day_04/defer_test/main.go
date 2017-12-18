package main

import "fmt"

func main()  {
	//a()
	b()
}
func a()  {
	//defer语句声明的变量，在其声明时就决定了
	i := 0
	defer fmt.Println(i)
	i ++
	return
}
func b(){
	//多个defer语句，后进先出
	for i := 0; i < 5;i++ {
		defer fmt.Println("%d", i)
	}
}
