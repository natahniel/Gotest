// @Time    : 2018/6/22 下午3:54
// @File    : 5数据类型.go

package main

import "fmt"

// 布尔类型,整型,浮点型
func varTest0() {
	// 布尔类型
	v1 := true // 布尔类型
	v2 := 999  // 整型
	v3 := 3.1  // float64
	fmt.Println("布尔类型true\t整型999\t浮点型3.1")
	fmt.Printf("布尔类型%T\t整型%T\t浮点型%T\n", v1, v2, v3)
}

// 字符类型,任意字符存到计算机中都是一个数字对应
func varTest1() {
	var chr0 byte // 声明字符类型
	chr0 = 97
	// 格式化输出,%c相应Unicode码点所表示的字符, %d十进制表示, %T相应值的类型的Go语法表示
	fmt.Printf("chr0 = %c and nunber is %d and type is %T\n", chr0, chr0, chr0)

	chr1 := '宇' // 自动推导,单引号的数据类型
	// 格式化输出,%c相应Unicode码点所表示的字符, %d十进制表示, %T相应值的类型的Go语法表示
	fmt.Printf("chr1 = %c and nunber is %d and type is %T\n", chr1, chr1, chr1)

	// 输出二进制表示
	fmt.Printf("朱宇杰的二进制表示为: %b/%b/%b\n", '朱', '宇', '杰')

	// 切换大小写
	fmt.Printf("A的小写%c\na的大写%c\n", 'A'+32, 'a'-32)

	// 字符串(字符类型只有一个,字符串可以有多个)
	chr2 := "朱宇杰"
	fmt.Printf("chr2 = %s and type is %T\n", chr2, chr2)
}

// 讨论字符和字符串的区别,字符串的切片和反引号
func varTest2() {
	chr0 := 'a'              // 字符都是单引号且一个(除转义符,多了个反斜杠)
	str0 := "明人不说暗话暗话,我是你爸爸" // 字符串是双引号且可以有多个,通常结尾默认隐藏转义符  \0
	fmt.Printf("字符 %c\n字符串 %s\n", chr0, str0)

	////切片问题,日后讨论
	////fmt.Println("字符串的切片[-1]: ", str0[-1])
	//fmt.Printf("字符串的切片[0]: %c\n", str0[0])
	//fmt.Printf("字符串的切片[2:3]: %s\n", str0[2:3])
	//fmt.Printf("字符串的切片[2:3]: %c\n", "hello world"[0])
	//fmt.Printf("字符串的切片[2:3]: %s\n", "hello world"[2:5])
}

//复数类型
func varTest3() {
	var com0 complex128 //定义一个复数,没有复制,会有一个初始化的值
	fmt.Printf("com0 = %v and type is %T\n", com0, com0)

	com0 = 9 + 7i //重新赋值
	//通过内建函数取出实部或虚部
	fmt.Printf("com0的实部 = %+v(%T), com0的虚部 = %+v(%T)\n", real(com0), real(com0), imag(com0), imag(com0))
}

//格式化占位符
func varTest4() {
	v0 := 2     //%d 匹配整型
	v1 := 3.1   //%f 匹配浮点型
	v2 := '我'   //%c 匹配字符
	v3 := "未于木" //%s 匹配字符串
	//格式化输出
	fmt.Printf("%d  %f  %c  %s\n", v0, v1, v2, v3)
	// %v 智能匹配,但会把字符输出为数字
	fmt.Printf("%v  %v  %v  %v\n", v0, v1, v2, v3)

}

// 入口函数,没有逻辑代码
func main() {
	//varTest0() // 布尔类型,整型,浮点型
	//varTest1() // 字符类型
	//varTest2() // 字符和字符串的区别
	//varTest3() //复数类型
	varTest4() //格式化占位符
}
