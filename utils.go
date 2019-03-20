package utils

import (
	"unsafe"
)

// StringToBytes 将string强制转换为[]byte
func StringToBytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

// BytesToString 将[]byte强制转换为string
func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
