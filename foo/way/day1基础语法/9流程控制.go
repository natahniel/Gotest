// @Time    : 2018/6/24 下午6:06
// @File    : 9流程控制.go

package main

import (
	"fmt"
	"math/rand"
	"time"
)

//Go语言支持最基本的三种程序运行结构：顺序结构、选择结构、循环结构。
//顺序结构：程序按顺序执行，不发生跳转。
//选择结构：依据是否满足条件，有选择的执行相应功能。
//循环结构：依据条件是否满足，循环多次执行某段代码。

func test900() (a int) {
	return 10
}

//选择结构 if 语句
func test90() {
	var myBalance int64
	fmt.Printf("请输入您的月薪: ")
	fmt.Scan(&myBalance)

	//判断月薪区间,报以不同态度
	if myBalance < 1000 { //月薪不足一千
		fmt.Println("您老打哪来回哪去吧")
	} else if myBalance < 2000 { //月薪不足两千
		fmt.Println("小伙子想吃点啥,土豆丝炒大白菜")
	} else if myBalance < 10000 { //月薪不到一万
		fmt.Println("哎呦,您里边请")
	} else if myBalance < 1000000 {
		fmt.Println("大爷,常来玩啊~")
	} else { //月入百万以上,土豪呀,超级无敌大土豪
		fmt.Println("您随意,您随意~")
	}

	//if支持一个初始化语句,初始化语句和判断条件用一个分号隔开 // 值得注意的是,此种方式初始化的变量在离开if代码块后找不到
	if var0 := test900(); var0 == 10 {
		fmt.Println("var0 = ", var0)
	} else {
		fmt.Println("var0 = ", var0)
	}

	var myDeposit int64         //用户存款
	fmt.Printf("请输入您的存款: ")     //提醒用户输入
	fmt.Scanf("%d", &myDeposit) //接受键盘输入并赋值给 myDeposit //用户存款

	//switch分支
	switch {
	case myDeposit < 10000:
		fmt.Println("贫困")
		break
	case myDeposit < 20000:
		fmt.Println("贫穷")
		break
	case myDeposit < 30000:
		fmt.Println("富裕")
		break
	case myDeposit < 40000:
		fmt.Println("小康")
		break
	default:
		fmt.Println("土豪")
		break
	}

}

//for循环 // go语言只有一个 for 循环,没有其他的诸如 while 循环
func test91() {
	sum := 0
	for i := 1; i <= 900000000; i++ {
		sum += i
	}
	fmt.Println(sum)

	name := "nathaniel"
	//range迭代返回两个值, 第一个为下标,第二个为值
	for _, i := range name {
		fmt.Printf("%c", i)
	}
}

// 跳转语句 break 和 continue
func test92() {
	currentFloor := 0 //勇士当前在一楼
	for ; currentFloor < 10; currentFloor++ {

		var0 := rand.New(rand.NewSource(time.Now().Unix())).Intn(10)
		time.Sleep(time.Second)

		if var0 > 7 { //找到公主了就赶紧溜
			fmt.Println(var0, "    ", currentFloor, "层, 找到公主了,带回家生猴子")
			break
		} else {
			if currentFloor == 9 {
				fmt.Println(var0, "    ", "最后一层也没找到公主,回家找五指姑娘吧")
			} else if var0 < 5 { //有可能遇到怪兽,打不过只能逃跑(怪兽层一定没有公主
				fmt.Println(var0, "    ", currentFloor, "层, 遇到怪兽了,打不过,跳过此楼层")
				continue
			} else {
				fmt.Println(var0, "    ", currentFloor, "层, 没找到公主,再上一层楼")
			}
		}
	}
}

func main() {
	//test90() //选择结构 if 语句, switch分支
	//test91() //for循环
	test92() //跳转语句 break 和 continue
}
