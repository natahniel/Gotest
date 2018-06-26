// @Time    : 2018/6/22 下午8:24
// @File    : 6等待键盘输入.go

package main

import "fmt"

//等待键盘输入
func test61() {
	var varInp int //声明变量
	varInp1 := "a" //自动推导变量类型

	fmt.Printf("请输入变量varInp>>")                                  //提示用户输入,不能像python一样 input(">>>")?
	fmt.Scanf("%d", &varInp)                                     //阻塞等待用户输入, 使用 & 表示键盘输入赋值给某变量
	fmt.Printf("varInp = %d and type is %T\n\n", varInp, varInp) //打印输入的变量

	fmt.Printf("请输入变量varInp1>>") //提示用户输入
	fmt.Scan(&varInp1)           //相当于 fmt.Scanf("%v", &varInp1)
	fmt.Printf("varInp1 = %s and type is %T\n", varInp1, varInp1)

}

func main() {
	test61()
}
