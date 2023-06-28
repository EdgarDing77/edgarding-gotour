package variable

/*
############################### 变量类型 ###############################
Go 的数据类型分类：

1. 布尔类型 bool
2. 数字类型 int、float32、float64等。并且支持复数，其中位的运算采用补码
3. 字符串类型：固定长度的字符连接起来的字符序列。go的字符串由单个字节连接起来，go的字符串字节采用UTF8编码标识 Unicode 文本
4. 派生类型：
    1. 指针 pointer
    2. 数组
    3. 结构化 struct
    4. channel
    5. 函数 func
    6. 切片
    7. 接口 interface
    8. Map

数字类型：（后面的数字分别对应占用的bit数）

- uint8、uint16、uint32、uint64
- int8、int16、int32、int64

浮点类型：

- float32、float54
- complex64（32位实数和虚数）、complex128（64位实数和虚数）

其他类型：

- byte （类似uint8）
- rune（类似int32）
- uint、int（32位或64位）
- uintptr（无符号整数，用于存放一个指针）
*/

// 基本类型
// 字符串
// 底层是 byte[] 使用 UTF-8 编码表示 Unicode 文本
// 两种表示形式：双引号识别转义字符/反引号原生输出
var varStr string = ""

// 整型
var varInteger int = 0 // int8、int16、int32、int64、uint、uint8、uint16、uint32、uint64、 uintptr

// 字符类型
// uint8 的别名，对应ascii码
var varByte byte = 0 // 'A'/65/'\x41'
// int32 的别名，对应UTF-8编码
var varRune rune = 0

// 浮点数类型
var varFloat float32 = 0 // float32、float64

// 复数类型
// 实际上由两个实数（在计算机中用浮点数表示，complex64对应的是float32，complex128对应的是 float64）构成，一个表示实部（real），一个表示虚部（imag）
var varComplex complex128 = complex(1, 2) // complex64、complex128

// bool型
// 只能存true/false，不支持自动或强制转换
// 优先级 ! > && > || 注意短路运算符
var varBool bool = false

var ptr *int
var arr []int
var set map[string]int
var channel chan int
var funcType func(string) int
var err error // error 是接口
