package variable

// 单行导入
import "fmt"

// Author: Vast Gui
// Email: guihy2019@gmail.com
// Date: 2022/11/6 11:21:41
// Description:
// Documentation:
// Reference:

// 这种因式分解关键字的写法一般用于声明全局变量
var (
	global_a int
	global_b bool
)

// Variable
//
// @Description: Go语言变量
func Variable() {
	// Go 语言变量名由字母、数字、下划线组成，其中首个字符不能为数字
	// 你可以在声明时就表明变量类型，当然也可以让其自动推断

	// 变量被声明后就必须被使用，但是全局变量可以不被使用

	// 声明字符串变量
	var a = "Runoob"
	fmt.Println(a)

	// 声明多个int类型变量
	var b, c = 1, 2
	fmt.Println(b, c)

	// 没有初始化就为零值
	var d int
	fmt.Println(d)

	// bool 零值为 false
	var e bool
	fmt.Println(e)

	// 数值类型（包括complex64/128）为 0
	// 布尔类型为 false
	// 字符串为 ""（空字符串）
	// 以下几种类型为 nil：
	// var a *int
	// var a []int
	// var a map[string] int
	// var a chan int
	// var a func(string) int
	// var a error // error 是接口
	var f string
	fmt.Println(f + "123")

	// 你也可以使用 := 来声明一个新的变量
	// 这是使用变量的首选形式，但是它只能
	// 被用在函数体内，而不可以用于全局变
	// 量的声明与赋值。使用操作符 := 可以
	// 高效地创建一个新的变量，称之为初始化声明
	g := 1
	fmt.Println(g)

	// 可以通过下列方式来快速交换两个变量的值
	// 但注意两个变量的类型必须一样
	h, i := 1, 2 // 声明变量
	h, i = i, h

	// 可以使用下列方式抛弃一个值
	_, i = i, h // 其中 i 的值被抛弃

	variable()
}

// 该函数小写命名，不能在包外被使用
func variable() {
	println("Hello")
}
