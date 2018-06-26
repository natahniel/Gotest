// @Time    : 2018/6/25 下午4:52
// @File    : testFile.go

package main

import (
	"math/rand"
	"time"
)

func test() (a int) {
	if var0 := rand.New(rand.NewSource(time.Now().Unix())).Intn(10); var0 > 5 {
		return 9
	} else {
		return 0
	}
}

func main() {
	test()
}
