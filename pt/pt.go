package pt

import (
	"fmt"
	"runtime"
)

// Ptf 普通格式化打印
var Ptf = func(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}

// Ptn 普通打印带换行
var Ptn = func(a ...interface{}) {
	fmt.Println(a...)
}

// Pt 普通打印
var Pt = func(a ...interface{}) {
	fmt.Print(a...)
}

// Ptlf 普通格式化打印带行号
var Ptlf = func(format string, a ...interface{}) {
	_, _, line, _ := runtime.Caller(1)
	fmt.Printf("l(%d): ", line)
	fmt.Printf(format, a...)
}

// Ptln 普通打印带行号带换行
var Ptln = func(a ...interface{}) {
	_, _, line, _ := runtime.Caller(1)
	if len(a) != 0 {
		fmt.Printf("l(%d): ", line)
	}

	fmt.Println(a...)
}

// Ptl 普通打印带行号
var Ptl = func(a ...interface{}) {
	_, _, line, _ := runtime.Caller(1)
	fmt.Printf("l(%d): ", line)
	fmt.Print(a...)
}
