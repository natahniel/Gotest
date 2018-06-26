// @Time    : 2018/6/23 下午7:48
// @File    : 8运算符.go

package main

import "fmt"

//算数运算符
func test80() { // 值得注意的是, go 的除法,两个整型相除, 结果亦为整型
	var0 := 99
	var1 := 88
	fmt.Printf("var0 + var1 = %d\nvar0 - var1 = %d\nvar0 * var1 = %d\nvar0 / var1 = %f\nvar0 %% var1 = %d\n",
		var0+var1, var0-var1, var0*var1, float64(var0)/float64(var1), var0%var1)
	fmt.Printf("%T", 9/3.3)
}

//关系运算符
func test81() { // 值得注意的是 = 表示赋值, == 表示相等
	fmt.Printf("3 > 2: %t\n3 < 2: %t\n3 == 2: %t\n3 <= 3: %t\n3 != 2: %t\n",
		3 > 2, 3 < 2, 3 == 2, 3 <= 3, 3 != 2)
}

//逻辑运算符
func test82() {
	var0 := false && true    // && 表示且
	var1 := false || true    // || 表示或
	var2 := !(true || false) // ! 表示非
	fmt.Printf("false && true = %t\nfalse || true = %t\n! (true || false) = %t\n", var0, var1, var2)
}

//赋值运算符
func test83() {
	var0 := 2
	var1 := 233
	var0 += var1
	fmt.Println("var0 += var1 相当于 var0 = var0 + var1, var0 = ", var0)
}

//位运算符
func test84() {
	var0 := 60 & 1 //按位与, var 按位与  1 可以判断奇偶性
	fmt.Println(var0)
}

func main() {
	//test80() //算数运算符
	//test81() //关系运算符
	//test82() //逻辑运算符
	//test83() //赋值运算符
	test84() //位运算符
}
