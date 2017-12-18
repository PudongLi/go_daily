package main
/*
	切片的基本用法
 */
import (
	"fmt"
)

func main() {
	sliceCopy()
}

func Sum(a[6] int) int  {
	count := 0
	for i := 0;i<len(a) ;i++  {
		count += a[i]
	}
	return count
	
}
func test2()  {
	a := [6]int{1,2,3,4,5,6}
	count := Sum(a)
	fmt.Println(count)
	b := a[2:3]
	c := a[:]
	c =append(c,b...)//将一个切片添加到另一个切片
	d := make([] int, 2, 5)
	fmt.Println(d)
}

func test3()  {
	a := make([]int, 5, 10)
	a[4] = 100
	b := a[0:10]
	b[9] = 100
	fmt.Printf("a=%#v len(a)=%d, cap(a)=%d\n", a, len(a), cap(a))
	fmt.Printf("b=%#v,len(b)=%d, cap(b)=%d\n", b, len(b), cap(b))
}

func sliceCopy(){
	a := []int{1,2,3,4}
	b := []int{5,5,5,5,5,5,5,5}
	//copy(b, a)
	//fmt.Printf("a=%#v, b=%#v\n", a, b)
	copy(a, b)
	fmt.Printf("a=%#v, b=%#v\n", a, b)

	}


func stringChange(){
	var str = "hello world"
	b := []byte(str)
	b[0] = 'a'
	str1 := string(b)
	fmt.Printf("str %s\n", str)
	fmt.Printf("str1 %s\n", str1)

}
func stringReversion(str string){

	b := []rune(str)
	length := len(b)
	for i:=0;i<length/2;i++{
		b[i] , b[length-i-1] = b[length-i-1], b[i]
	}
	str1 := string(b)
	fmt.Println(str1)
}
func countHanzi(str string){
	b := []rune(str)
	fmt.Printf("length:%d\n", len(b))
}
