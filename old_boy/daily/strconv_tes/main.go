package main


import (
	"fmt"
	"strconv"
	"reflect"
)

func main() {
	StringToBool()
	BoolToString()
	StringToFloat()
	StringToInt()
	StringToInt2()
}

func StringToInt2(){
	str := "123456"
	num, _ := strconv.Atoi(str)
	fmt.Println(reflect.TypeOf(num))
}

func StringToInt(){
	str := "12345678"
	num, _:= strconv.ParseInt(str,10, 64)
	fmt.Println(reflect.TypeOf(num))
}
func StringToBool(){
	str := "True"
	result, _:= strconv.ParseBool(str)
	fmt.Println(result)
}
func BoolToString()  {
	result := strconv.FormatBool(1 > 0)
	fmt.Println(result)
	
}

func StringToFloat ()  {
	str := "12.12"
	result , _:= strconv.ParseFloat(str, 64)
	strconv.ParseFloat()
	fmt.Println(reflect.TypeOf(result))
}
