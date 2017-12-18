package main

import "fmt"

const (
	a = iota
	b
	c
	d = 1
	f = 2
	g
	h
	i = iota
	j
	k
)

func main() {

	fmt.Printf("%d %d %d %d %d %d %d %d %d %d\n",a, b, c, d, f, g, h, i, j, k)
}