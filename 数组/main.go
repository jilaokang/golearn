package main

import (
	"fmt"
)

func main() {
	var a [2]string

	a[0] = "hello"
	a[1] = "world"

	// helo world
	fmt.Println(a[0], a[1])
	// [hello world]
	fmt.Println(a)
}
