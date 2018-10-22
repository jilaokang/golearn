package main

import (
	"fmt"
)

type Vertex struct {
	x int
	y int
}

func main() {
	// 输出结果为 {1 2}
	v := Vertex{1, 2}
	v.x = 12

	fmt.Println(v)

	p := &v
	p.x = 1e9
	fmt.Println(v)
}

// 总结：结构体类似于js里的对象，创建对象并赋值，使用对象。
//      可以通过指针间接访问
