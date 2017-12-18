package main

import "fmt"

/*
函数是一等公民，可以像变量一样赋值
 */
type MyInt int

type add_func func(int, int) int

func add(a, b int) int {
	return a + b
}

func operator(op add_func, a int, b int) int {
	return op(a, b)
}

func testType(){
	var a MyInt
	a = 100
	a += 100
}
func test1(){
	c := add
	fmt.Printf("%p %T %p %T\n",c, c, add, add)
	sum := c(10, 20)
	fmt.Println(sum)
	sum = add(10, 20)
	fmt.Println(sum)
}
func test2()  {
	c := add
	fmt.Println(c)
	sum := operator(c, 100, 200)
	fmt.Println(sum)
}
func test3(a *int){
		*a = 100
}
func test4(a, b int) (c int){
	c = a + b
	return
}
func Add(arg...int) int{
	var sum int
	for i := 0; i < len(arg); i++ {
		sum = sum + arg[i]
	}
	return sum
}
func Add2(a int,arg...int) int{
	var sum int
	for i := 0; i < len(arg); i++ {
		sum = sum + arg[i]
	}
	return sum + a
}

func main() {
	//test1()
	//test2()
	//a :=1
	//fmt.Println(a)
	//test3(&a)
	//fmt.Println(a)
	//result := test4(1,2)
	//fmt.Println(result)
	fmt.Println(Add(1,2,3,4,5))
	fmt.Println(Add2(1,2))
}