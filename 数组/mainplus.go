package main

import (
	"fmt"
)

func main() {
	p := []int{1, 2, 3, 4, 5}

	// [12,3,4,5]
	fmt.Println(p)

	for i := 0; i < len(p); i++ {

		fmt.Println("p[%d] == %d\n", i, p[i])
	}

	// [2 3] 表示从起点到（尾值-1】区间
	fmt.Println(p[1:3])
	// 【2, 3, 4, 5】从1到结束
	fmt.Println(p[1:])
	// [1,2,3] 从0到3
	fmt.Println(p[:3])

	// 类似push,第一个为数组，后面为push值
	p = append(p, 2, 3, 4)
	fmt.Println(p)

	for i, v := range p {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
