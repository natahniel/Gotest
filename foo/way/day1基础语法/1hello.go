/*
// go语言以包作为管理单位, 每个文件必须先声明包,包是GO语言里最基本的分发单位,也是工程管理中依赖关系的体现
// 程序要运行必须要有main包
*/
package main

// 导入包后必须使用,否则报错
import "fmt" // 需要使用fmt包中的Println()函数

// 入口函数,不能带参数也不能定义返回值
func main() { //左括号不可另起新行
	// 一段一段处理,结尾自动加换行
	fmt.Println("hello world.你好,世界")

	testVar := 998
	// 格式化输出,结尾不会自动换行
	fmt.Printf("%d的数据类型是%T\n", testVar, testVar)
}
