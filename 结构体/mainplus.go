package main

import (
	"fmt"
)

type Vartex struct {
	X, Y int
}

var (
	// {1 1}
	v1 = Vartex{1, 1}
	// {2 0}
	v2 = Vartex{X: 2}
	// 0 0
	v3 = Vartex{}

	//&{4 4}
	p = &Vartex{4, 4}
)

func main() {
	fmt.Println(v1, v2, v3, p)
}
