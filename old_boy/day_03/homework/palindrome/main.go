package main
/*
给定一个字符串 判断其是否是回文
 */
import "fmt"

func main() {
	str := "上海自来水来自海上"
	result := palindrome(str)
	if result{
		fmt.Printf("%s  是回文", str)
		return
	}
	fmt.Printf("%s  不是回文", str)
}

func palindrome(str string) bool{
	strSlice := []rune(str)
	for i := 0;i < len(strSlice)/2;i++ {
		if strSlice[i] != strSlice[len(strSlice)-i-1]{
			return false
		}
	}
	return true
}