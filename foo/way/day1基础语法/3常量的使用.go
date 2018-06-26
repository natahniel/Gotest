package main

import "fmt"

func main() {
	// 变量: 是程序运行期间可以改变的量  var 关键字
	// 常量: 是程序运行期间不可改变的量  const 关键字
	const a int = 9
	// a = 10 // cannot assign to a,不能给常量赋值
	fmt.Println(a)

	//常量也可以自动推导
	const b = "朱宇杰"
	fmt.Printf("b type is %T\n", b)

	// 定义多个变量并自动推导类型
	var (
		v1 = 9
		v2 = "哈哈"
	)
	// 定义多个常量并自动推导类型
	const (
		v3 = "10"
		v4 = "11"
	)
	fmt.Printf("变量v1 = %d; 变量v2 = %s; 常量v3 = %s; 常量v4 = %s\n", v1, v2, v3, v4)
}
