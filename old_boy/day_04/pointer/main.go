package main

import "fmt"

func main() {
	var b int  = 1
	var p *int
	p = &b
	modify(p)
	fmt.Println(b)

}
func modify(a *int){
	*a =100
}
func test1(){
	var a int
	a = 10
	var b *int
	fmt.Println(a)
	fmt.Printf("%p\n",b)//输出b的值
	fmt.Println(&b)//输出b的地址

	b = &a
	fmt.Println(*b)
	*b = 100
	fmt.Println(*b)
	fmt.Println(a)
}