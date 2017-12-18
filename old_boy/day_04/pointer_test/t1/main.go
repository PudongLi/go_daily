package main

import "fmt"

func main() {
	//getPointer()
	//test2()
	test4()


}
func test4(){
	var a []int
	a = make([]int,10)
	a[0] = 100
	fmt.Println(a)

	var p *[]int
	p = new([]int)
	*p = make([]int,10)
	(*p)[0] =100
	fmt.Println(p)
	p = &a
	(*p)[0] = 1000
	fmt.Println(a)

}
func test3()  {
	var p *int
	var b int
	p = &b
	*p = 200
	p = new(int)
	*p = 1000
}

func test2(){
	var a = 2
	var p *int
	p = &a
	changeValue(p)
	fmt.Println(a)
}
func getPointer(){
	//获取变量地址并打印到终端
	var a = 10
	fmt.Printf("%p\n", &a)
}
func changeValue( b *int){
	//传入一个int类型的指针，在函数中修改所指向的值
	*b = 100

}
