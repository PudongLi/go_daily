package main

import (
	"fmt"
	"math/rand"
	"time"
)
var (
	randSource = rand.New(rand.NewSource(time.Now().Unix()))

)

func createNumberList()[]int{
	var NumberList = []int {}
	for i := 0; i < 10; i++ {
		NumberList = append(NumberList, randSource.Intn(100))
	}
	return  NumberList
}

func main() {
	NumberList := createNumberList()
	fmt.Println(NumberList)

	bList := BubbleSort(NumberList)
	fmt.Println(bList)

	sList := SelectionSort(NumberList)
	fmt.Println(sList)
	qList := QuickSort(NumberList, 0, 9)
	fmt.Println(qList)


}

func BubbleSort(bList []int) []int{
//每次比较两个相邻的元素，如果他们的顺序错误就把他们交换过来
//复杂度 O(n^2)

	for i := 0; i < len(bList); i++ {

		for j := 1; j < len(bList) - i; j++ {
			if bList[j] < bList[j-1]{
				bList[j], bList[j-1] = bList[j-1], bList[j]
			}

		}
	}
	return bList

}

func QuickSort(qList []int, left int, right int) []int{
	//分治法

	if left < right{
		var i = left
		var j = right
		var x = qList[left]

		for i < j{
			for (i < j) && (qList[j] >= x){
				j --
			}
			if i < j{
				i++
				qList[i] = qList[j]
			}
			for (i < j) && (qList[i] < x){
				i ++
			}
			if i < j{
				j--
				qList[j] = qList[i]
			}
		}

		qList[i] = x
		QuickSort(qList, left, i-1)
		QuickSort(qList,i+1, right)
	}
	return qList
}

func SelectionSort(sList []int) []int{
	bLen := len(sList)
	for i := 0; i < bLen; i++ {
		max := 0
		for j := 1; j < bLen - i; j++ {
			if sList[j] > sList[max]{
				max = j
			}
		}
		sList[bLen-i-1], sList[max] = sList[max], sList[bLen-i-1]
	}
	return sList
	}

func InsertionSort(iList []int) []int{

	iLen := len(iList)

	for i := 1; i < iLen; i++ {
		if iList[i] < iList[i-1]{
			j := i - 1
			temp := iList[i]
			for j >= 0 && iList[j] > temp{
				iList[j+1] = iList[j]
				j --
			}
			iList[j+1] = temp
		}
	}

}

