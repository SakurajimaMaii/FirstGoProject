package variable

import "fmt"

// Author: Vast Gui
// Email: guihy2019@gmail.com
// Date: 2022/11/6 11:21:41
// Description:
// Documentation:
// Reference:

// Variable
//
//	@Description: Go语言变量
func Variable() {
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
}
