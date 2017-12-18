package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var numberList1 [10] int
	var numberList2 [10] int
	var numberList3 [10] float64
	n := rand.New(rand.NewSource(time.Now().Unix()))

	for i := 0 ; i < 10; i ++ {
		numberList1[i] = n.Int()
	}

	for j := 0 ; j < 10; j ++ {
		numberList2[j] = n.Intn(100)
	}

	for k := 0 ; k < 10; k ++ {
		numberList3[k] = n.Float64()
	}

	fmt.Printf("10个随机整数:%v\n", numberList1)
	fmt.Printf("10个小于100的随机数:%v\n", numberList2)
	fmt.Printf("10个随机浮点数:%v\n", numberList3)

}

