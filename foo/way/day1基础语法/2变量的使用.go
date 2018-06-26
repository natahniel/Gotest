package main

import "fmt"

// 入口函数
func main(){
    // 变量赋值前必须声明
    // a = 99 // 报错 变量没有声明

    // 同一个{}里声明的变量是唯一的,不可再次申明同名变量
    var a int // 声明一个变量 int 类型
    // var a strint // 报错
    a = 999 // 给变量赋值,可以重复赋值相当于修改
    // 声明变量后必须使用,否则报错
    fmt.Println("hello", a)

    // go是强类型语言,变量数据类型一旦确定不可修改
    // name = 999 //执行,将发生错误,不能使用int类型给string类型的变量赋值
    // fmt.Println(name) // cannot use 999 (type int) as type string in assignment

    // 以此种方式声明的变量无需 var 关键字,且会自动推导数据类型
    name := "nathaniel" // 姓名
    height := 175 // 身高 cm
    weight := 60 // 体重 kg
    fmt.Printf("%s的身高是%d,体重是%d。\n", name, height, weight)

    // go语言中单引号和双引号表示的意义是不同的
    testVar := '哈'
    testVar1 := "123"
    fmt.Printf("go语言中,单引号的数据类型是%T,双引号的数据类型是%T\n", testVar, testVar1)

    // 多重赋值 // 先计算等号右边的,再赋值给等号左边
    applePrice, pearsPrice := 10, 20
    pearsPrice, applePrice = applePrice, pearsPrice
    fmt.Printf("applePrice:%d, pearsPrice:%d\n", applePrice, pearsPrice)

    // 匿名变量,丢弃数据不处理,常用于配合函数返回值使用
    var retVar int
    _, retVar, _ = testFun()
    fmt.Printf("retVar = %d\n", retVar)
}

func testFun()(a, b, c int){
    return 1, 2, 3
}









