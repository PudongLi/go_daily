package main

/*
输入一个字符串，分别统计其中英文字母、空格、数字和其他字符的个数
 */
import (
	"fmt"
)
func main() {
	count("123ello45ATest...")
}

func count(str string)  {
	strSlice := []byte(str)
	letter := 0 //字母
	number := 0 //数字
	other  := 0 //其他字符
	for i := 0;i <len(str);i++ {
		// 可直接判断A-Z a-z 0-9用字符串表示
		if(strSlice[i]>=65 && strSlice[i]<=90)||(strSlice[i]>=97 && strSlice[i]<=122){
			letter ++
			continue
		}
		if strSlice[i]>=48 && strSlice[i]<=59 {
			number ++
			continue
		}
		other ++
	}
	fmt.Printf("在字符串%s中，英文字母有%d个，数字有%d个，其他字符有%d个\n", str, letter, number, other)


}
