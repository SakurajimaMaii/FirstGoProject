package constant

// 常量通常是运行时不会被修改的数值，使用 constant 声明
// 常量依旧可以被声明但没有使用
const first_constant = 3.14

// 常量还可以被用作声明枚举变量，例如
const (
	Unknown = 0
	Female  = 1
	Male    = 2
)

// iota，特殊常量，可以认为是一个可以被编译器修改的常量。

// iota 在 const 关键字出现时将被重置为 0 (const 内部的第一行之前)，
// const 中每新增一行常量声明将使 iota 计数一次(iota 可理解为 const 语句块中的行索引)。

// iota 可以被用作枚举值：
const (
	a = iota
	b = iota
	c = iota
)

const (
	i = 1 << iota
	j = 3 << iota
	k
	l
)
