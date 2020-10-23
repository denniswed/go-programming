package main

import (
	"fmt"
)

type theanswertoeverything int

var x theanswertoeverything

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}
