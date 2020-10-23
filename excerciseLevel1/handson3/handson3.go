package main

import (
	"fmt"
)

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	fmt.Println(x, y, z)
	//zero value

	s := fmt.Sprintf("%T\t%v\t%T\t%v\t%T\t%v", x, x, y, y, z, z)
	fmt.Println(s)
}
