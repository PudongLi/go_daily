package main

import (
	"fmt"
	"math/rand"
	"time"
)
var(
	NewTimeSource = rand.New(rand.NewSource(time.Now().Unix()))
	str = "aadadfsjrmcmqwqw1243rsxvdgetsadasdewrfhphbpbfr"

)

func main() {

	//a := randomInt()
	//fmt.Printf("%#v\n", a)
	b := randomString()
	fmt.Printf("%#v\n", b)


}

func test1(){
	var a = [5] string {0:"age", 2:"name"}
	fmt.Printf("%#v\n", a) // %# 以go的方式打印
}

func randomInt() [100]int{
	var a = [100]int{}
	for i := 0;i < 100; i++ {
		a[i] = NewTimeSource.Int()

	}
	return  a
}

func randomString()[100]string{
	var a = [100]string{}
	for i := 0;i < 100; i++ {
		var s string
		for j := 0;j < 32;j++ {
			index := NewTimeSource.Intn(len(str))
			s  = fmt.Sprintf("%s%c", s, str[index])
			//str  = str + strArray[NewTimeSource.Intn(strLen)]
		}
		a[i] = s

	}
	return a
}
