package mem

import (
	"reflect"
	"unsafe"
)

// 移动指针
func MovePtr(p unsafe.Pointer, step uintptr) unsafe.Pointer {
	pU := uintptr(p) + step
	return unsafe.Pointer(pU)
}

// 仅仅为提示归纳，平时直接用refect 包
func MemCmp(src interface{}, des interface{}) bool {
	return reflect.DeepEqual(src, des)
}
