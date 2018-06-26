package main

import "fmt"

func main() {
	// iota常量自动生成器，每次自动累加1,给常量赋值使用
	const (
		v0 = iota
		v1 = iota
		v2 = iota
	)
	fmt.Printf("v0 = %d, v1 = %d, v2 = %d\n", v0, v1, v2)

	// iota遇到const重置为0
	const v00 = iota
	fmt.Printf("v00 = %d\n", v00)

	// 可以只写一个iota,如果在同一行,值相同
	const (
		a = iota
		b
		c
	)
	fmt.Printf("a = %d, b = %d, c = %d\n", a, b, c)

	// 可以只写一个iota,如果在同一行,值相同
	const (
		a1, b1, c1 = iota, iota, iota
	)
	fmt.Printf("a1 = %d, b1 = %d, c1 = %d\n", a1, b1, c1)

}
