package main

import (
	"fmt"
)

type theanswertoeverything int

var x theanswertoeverything
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)

	y = int(x)
	fmt.Println(x)
	fmt.Printf("%T\n", x)
}
