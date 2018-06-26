// @Time    : 2018/6/23 下午6:52
// @File    : 7类型转换和类型别名.go

package main

import "fmt"

//数据类型的转换
func test71() {
	////可以转换的类型称之为兼容类型,不可以转换的类型称之为不兼容类型
	//flag := true      // 布尔类型
	//int0 := int(flag) // 不能将布尔类型转换为 int 类型 // Cannot convert expression of type bool to type int

	//字符类型本质上就是整型
	var0 := 'o'
	var1 := int(var0)
	fmt.Printf("var1 = %d and type = %T\n", var1, var1)
	fmt.Printf("var1 = %c and type = %T\n", var1, var1)

	//var2 := "9"
	//var3 := int(var2) // 这种操作在python是允许的,但是golang中不允许  // cannot convert var2 (type string) to type int
}

//类型别名
func test72() {
	//给类型 string 起别名叫做 str
	type str string
	var var0 str
	fmt.Printf("var0 = %s and type is %T\n", var0, var0)
}

func main() {
	test71()
	test72()
}
