// @Time    : 2018/6/25 下午4:30
// @File    : 2递归函数.go

package main

import "fmt"

//斐波那切数列
func test20(n int) (sum int) {
	if n > 1 {
		sum = test20(n-1) + test20(n-2)
		return sum
	} else {
		return sum
	}
}

func main() {
	fmt.Println(test20(5))
}
