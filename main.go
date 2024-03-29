/*
 * MIT License
 *
 * Copyright (c) 2023 Vast Gui
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

// 包声明
package main

// 导入包
// Go 语言导入包有两种形式，单行导入和多行导入
// 下面展示的就是多行导入
import (
	"FirstGoProject/variable"
	"fmt"
)

// 函数
func main() {
	// 变量
	var name = "Hello world."
	fmt.Println(name)
	variable.Variable()
}

//当标识符（包括常量、变量、类型、函数名、结构字段等等）以一个大写字母开头，
//如：Group1，那么使用这种形式的标识符的对象就可以被外部包的代码所使用
//（客户端程序需要先导入这个包），这被称为导出（像面向对象语言中的 public）；
//标识符如果以小写字母开头，则对包外是不可见的，但是他们在整个包的内部是可
//见并且可用的（像面向对象语言中的 protected ）。
