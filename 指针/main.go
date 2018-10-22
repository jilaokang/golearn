package main

import (
	"fmt"
)

func main() {
	i, j := 42, 2701

	// p为指向i的指针
	p := &i
	// 读取指针对应的值 42
	fmt.Println(*p)
	*p = 21
	// 指针改变i 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	// 指针值改变 73
	fmt.Println(j)
}

// 总结：指针是直接指向对饮内存位置的东东，操控指针的变化是实质操控该内存所在变量的值
