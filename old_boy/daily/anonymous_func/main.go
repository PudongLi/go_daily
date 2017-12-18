package main

import (
	"fmt"
	"time"
)
var(
	ipList = []string{"1", "2", "3", "4"}

)
func main() {
	for _, ip := range ipList{
		go func(v string) {

			fmt.Println(v)
		}(ip)
	}
	time.Sleep(time.Second*5)
}

/*
func work(){
	//normal
	for _, ip := range ipList{
		go func(v string) {
			fmt.Println(ip)
		}(ip)
	}
	time.Sleep(time.Second*5)
}
func work2(){
	//abnormal
	for _, ip := range ipList{
		go func() {
			fmt.Println(ip)
		}()
	}
	time.Sleep(time.Second*5)
}*/
