// @Time    : 2018/6/25 上午10:53
// @File    : 参数和返回值.go

package main

import "fmt"

//函数 无参无返回值
func test10() {
	fmt.Println("明人不说暗话,我是你爸爸")
}

//函数 有参无返回值
func test11(var0, var1 int, name string) {
	fmt.Printf("var0 + var1 = %d\n", var0+var1)
	fmt.Printf("我的名字叫%s\n", name)
}

//不定长参数,无返回值
func test12(args ...int) {
	sum := 0
	for _, i := range args {
		sum += i
	}
	fmt.Println(sum)
}

//无参有返回值
func test13() (a, b, c int) {
	a = 998
	return a, 2, 3
}

//有参有返回值
func test14(name string, age int) (theClass string) {
	fmt.Printf("姓名: %s, 年龄: %d, 班级: 三年级六班\n", name, age)
	return "三年级六班"
}

func main() {
	//test10() //无参无返回值函数调用,没有默认返回值
	//test11(2, 4, "朱宇杰") //有参无返回值函数调用
	//test12(1, 2, 3, 4, 5, 699, 123) //不定长参数,无返回值
	//fmt.Println(test13())        //无参有返回值函数调用
	//var0, var1, var2 := test13() //可以将返回值赋值给变量
	//fmt.Printf("var0 = %d, var1 = %d, var2 = %d\n", var0, var1, var2)
	test14("nathaniel", 18) //有参有返回值
}
