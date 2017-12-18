package main

import "fmt"

/*
计算两个大数相加之和
 */
func main() {
	num1 := "11111111111111111111111111111111111"
	num2 := "1111111111111111111111111111111111111111111111111111111111"
	//num1 := "15005"
	//num2 := "5005"
	num3 := "5555"
	num4 := "6666"
	add(num1, num2)
	add(num3, num4)
}

func add(num1 string, num2 string)  {
	num1Byte := []byte(num1)
	num2Byte := []byte(num2)
	result := []int{}
	carryFlag := 0 //是否进位
	if len(num2Byte) > len(num1Byte){
		num1Byte, num2Byte = num2Byte, num1Byte
	}
	len1 := len(num1Byte)
	len2 := len(num2Byte)
	//短的补零
	j := len2 - 1
	for i := len1 - 1; i >= 0; i-- {
		presentSum := 0
		if j >= 0 {
			presentSum = int(num1Byte[i]) + int(num2Byte[j]) + carryFlag - 48 * 2
		}else{
			presentSum = int(num1Byte[i]) + carryFlag - 48
		}
		if presentSum >= 10 {
			presentSum = presentSum - 10
			carryFlag = 1
		}else {
			carryFlag = 0
		}

		result = append(result, presentSum)
		j --
	}
	if carryFlag == 1{
		result = append(result, 1)
	}
	result2 := stringReversion(result)
	fmt.Println("结果：")
	for _, r := range result2{
		fmt.Printf("%d", r)
	}
	fmt.Println()
}

func stringReversion(b []int) [] int{
/*
	反转
 */
	length := len(b)
	for i:=0;i<length/2;i++{
		b[i] , b[length-i-1] = b[length-i-1], b[i]
	}
	return b
}
