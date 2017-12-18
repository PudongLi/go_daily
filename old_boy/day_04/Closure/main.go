package main
/*
闭包
 */
import (
	"fmt"
	"strings"
)

func main() {

}
func testmakeSuffixFunc(){
	func1 := makeSuffixFunc(".bmp")
	func2 := makeSuffixFunc(".jpg")
	fmt.Println(func1("test"))
	fmt.Println(func2("test"))
}
func testAdder(){
	var add = Adder()
	fmt.Println(add(1))
	fmt.Println(add(20))
	fmt.Println(add(300))
}

func Adder() func(i int) int{
	var x int
	return func(j int) int {
		x += j
		return x
	}
}

func makeSuffixFunc(suffix string)func(string) string{
	return func(name string) string {
		if !strings.HasSuffix(name, suffix){
			return name + suffix
		}
		return name
	}
}