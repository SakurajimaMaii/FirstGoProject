package main

import (
	"FirstGoProject/variable"
	"fmt"
)

// Hello World函数
func main() {
	var a = 123
	var b = "2022-11-6"
	var format = "age=%d,date=%s"
	var result = fmt.Sprintf(format, a, b)
	fmt.Println(result)
	variable.Variable()
}
