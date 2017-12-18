package main

import (
	"fmt"
)

func main() {
	ArrayName1 := [...]string{"lipd", "wangyiru", "hahaha"}
	ArrayName2 := ArrayName1

	fmt.Printf("%v %v \n", ArrayName1[1], ArrayName2[1])
	ArrayName1[1] = "zhuzhuxia"
	fmt.Printf("%v %v \n", ArrayName1[1], ArrayName2[1])
	fmt.Printf("%v,%v %v \n", &ArrayName1[0], &ArrayName1[1], &ArrayName1[2])

}
