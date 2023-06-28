package variable

/*
############################### 常量 ###############################
*/

// 常量定义方式
// const := 不允许这样使用
const a, b, c = 1, 2, true
const (
	v1         = "value1"
	v2         = "value2"
	v3 int     = 3
	v4 float32 = 4.0
)

// 枚举
// 可用iota关键字简化：每一行相比与上一行iota的值都会默认加1
// iota 特殊常量
const (
	var1 = iota*iota + 1 // 0*0 + 1 1
	var2                 // 1*1 + 1 2
	var3                 // 2*2 + 1 5
	var4 = 100           // 100 iota=3
	var5                 // 100  iota=4
	var6 = iota          // iota=5
)
