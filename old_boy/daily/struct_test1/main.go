package main
/*
结构体指针

 */
import (
	"fmt"
)

func main() {
	var rect1 = new(Rect)
	//var rect1 = Rect{20,30}
	rect1.length = 30
	rect1.width = 20
	fmt.Println(*rect1)
	fmt.Println(rect1.Area())
	fmt.Println(*rect1)
}

type Rect struct {
	width, length float64
}

func (rect *Rect)Area() float64  {
	rect.width += 200
	rect.length += 200

	return rect.width * rect.length
}
