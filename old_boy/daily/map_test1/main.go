package main
/*
map是键值对未排序的集合
 */
import (
	"fmt"
)

func main() {
	map1 := make(map[string]string)
	map1["name"] = "lipd"
	map1["age"]  = "24"
	map1["city"] = "beijing"

	var keys [] string
	for _, key2 := range map1{
		keys = append(keys, key2)
	}
	value, ok := map1["test"]
	if ok{
		fmt.Println(value,ok)
	}
}
