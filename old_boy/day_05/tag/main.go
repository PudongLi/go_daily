package main

/*
tag
 */
import (
	"fmt"
	"encoding/json"
)

type Student struct {
	Name string `json:"name"`
	Age int     `json:"age"`
}

func main() {
	var s Student
	s.Age = 200
	s.Name = "abc"
	//student的属性只能大写，小写的话其他的包里的函数访问不到这个属性
	data, _:= json.Marshal(s)//将结构体序列化为json字符串
	fmt.Println(string(data))
	var str Student
	_ = json.Unmarshal(data, &str)//反序列化
	fmt.Println(str)
	}


