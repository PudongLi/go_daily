package main

import "fmt"

func main() {


}
func test1(){
	map1 := make(map[int]int,10)
	map1[1] =1
	fmt.Println(map1)
	testMap(map1)
	fmt.Println(map1)
}
func sliceOfMap(){
	items:= make([]map[int]int, 5)
	for i := 0; i <= len(items); i++{
		items[i] = make(map[int]int,10)
	}
	items[0][1] = 10
	fmt.Println(items)
	//对map排序输出

}


func testMap(map1 map[int]int)  {
	map1[1] = 100
	//var a map[string]string = map[string]string{"hello", "world"}
	a := make(map[string]string,10)//必须make，为其分配内存
	a["hello"] = "world"  //插入，大于10时会自动扩容
	a["name"] = "lipd"
	for k, v := range a{
		fmt.Printf("%s %s\n", k, v)
	}
	val, ok := a["hello"] //查找，返回值和bool
	if ok {
		fmt.Println(val)
	}

	delete(a, "hello")//删除
	fmt.Println(a)
}