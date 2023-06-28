package variable

/*
############################### 变量定义 ###############################
1. 指定变量类型 `var identifier type`
2. 根据值自行判断变量类型 `var v_name = value`
3. `:=` 声明（与var声明冲突）`v_name := value`
*/

// var 指定变量
// 使用范围：不限
var globalVar string = "value"
var globalVar2 = "value"

func funcVar() string {
	var3 := "value"
	return var3
}
