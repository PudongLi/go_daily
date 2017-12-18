package main

import "fmt"

type Phone struct {
	price int
	color string
}

func (phone Phone)ringing()  {
	fmt.Println("dinglinnnnng")
}

type IPhone struct {
	Phone
	model string
}

func main() {
	var myPhone = new(IPhone)
	myPhone.model = "hhh"
	myPhone.price = 1000
	myPhone.color = "black"
	myPhone.ringing()

}