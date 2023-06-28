package flowcontrol

import (
	"fmt"
	"time"
)

/*
############################### 流程控制 ###############################
*/

func FlowControl() {
	// if
	a := 1
	if a < 10 {
		fmt.Println("第一种写法，此时a是函数体内的变量")
	} else if b := 1; b < 10 {
		fmt.Println("第二种写法，此时b是if中的变量")
	} else {
		fmt.Println("啥也不是")
	}

	// switch
	flag := 0
	switch flag {
	case 0, 1:
		fmt.Println("print a")
		fallthrough // 只会执行下一个case
	case 2:
		fmt.Println("already printed1") // 尽管条件语句判断为false，但是由于前面 fallthrough，此case会被强制执行
		if flag != 1 {
			fmt.Println("case is break")
			break
		}
		fmt.Println("already printed2")
	case 3:
		fmt.Println("not print yet") //不会执行
	default:
		fmt.Println("default")
	}

	switchType := func(a interface{}) {
		// switch type
		switch a.(type) {
		case bool:
			fmt.Println("flag is bool type")
		case int:
			fmt.Println("flag is int type")
		default:
			fmt.Println("flag is other type")
		}
	}
	switchType(flag)

	// select
	// 用来监听 channel 上的数据流动，多辅以for轮询
	// 若任意语句可以继续执行（即使没有被阻塞），*随机*选择一条语句来进行先
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		c1 <- "one"
		time.Sleep(1 * time.Second)
	}()
	go func() {
		c2 <- "two"
		time.Sleep(2 * time.Second)
	}()

	for i := 0; i < 4; i++ {
		time.Sleep(3 * time.Second)
		select {
		case msg1 := <-c1: // 每个case都必须是IO操作
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		default:
			// 如果两个通道都没有可用的数据，则执行这里的语句
			fmt.Println("no message received")
		}
	}
}
