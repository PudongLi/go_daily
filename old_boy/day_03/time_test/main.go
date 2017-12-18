package main

import (
	"fmt"
	"time"
)

func main() {
	calcSpendTime()
	}

func getNow()  {
	now := time.Now()
	fmt.Printf("%02d%02d%02d\n", now.Hour(),now.Minute(), now.Second())

}
func constTimeValue(){
	fmt.Printf("%d\n", time.Nanosecond)
	fmt.Printf("%d\n", time.Microsecond)
	fmt.Printf("%d\n", time.Millisecond)
	fmt.Printf("%d\n", time.Second)
	fmt.Printf("%d\n", time.Minute)
	fmt.Printf("%d\n", time.Hour)

}

func formatTime(){
	now := time.Now()
	str := now.Format("2006/01/02 15:04:05")
	fmt.Println(str)
}
func calcSpendTime()  {
	now1 := time.Now()
	fmt.Println(now1.Nanosecond())
	time.Sleep(time.Second*2)
    spend := time.Since(now1)
    str := int(spend)
    fmt.Println(str)
}