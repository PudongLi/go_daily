package main

import "fmt"

func main() {
	//var b bool
	//var a int8 = -127
	//fmt.Printf("%t\n", b)
	//fmt.Printf("%d\n", a)
	var user User
	user.name = "lipd"
	user.pre = 0
	fmt.Printf("user is DaRen:%v\n", GetPre(user,1<<1))
	user = SetPre(user, true, 1<<1)
	user = SetPre(user, true, 1<<0)
	fmt.Println(user.pre)
	fmt.Printf("user is DaRen:%v\n", GetPre(user,1<<1))
	//user = SetPre(user, true, 1<<1)
	//fmt.Printf("user is DaRen:%v\n", GetPre(user,1<<1))
	//user = SetPre(user, false, 1<<1)
	//fmt.Printf("user is DaRen:%v\n", GetPre(user,1<<1))


}
