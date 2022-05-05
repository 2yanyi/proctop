package stringx

import (
	"reflect"
	"unsafe"
)

/**
 * Provides low-level type conversion without generating a memory copy.
 *
 * 提供底层类型转换，不会产生内存拷贝。
 *
 */

// BytesToString is equivalent to string(a)
func BytesToString(a []byte) (_ string) {
	if a != nil {
		return *(*string)(unsafe.Pointer(&a))
	}
	return
}

// RunesToString is equivalent to string(a)
func RunesToString(a []rune) (_ string) {
	if a != nil {
		return *(*string)(unsafe.Pointer(&a))
	}
	return
}

// StringToBytes is equivalent to []byte(a)
func StringToBytes(a *string) (_ []byte) {
	if a != nil {
		return *(*[]byte)(unsafe.Pointer(str2slice((*reflect.StringHeader)(unsafe.Pointer(a)))))
	}
	return
}

// StringToRunes is equivalent to []rune(a)
func StringToRunes(a *string) (_ []rune) {
	if a != nil {
		return *(*[]rune)(unsafe.Pointer(str2slice((*reflect.StringHeader)(unsafe.Pointer(a)))))
	}
	return
}

func str2slice(str *reflect.StringHeader) *reflect.SliceHeader {
	return &reflect.SliceHeader{
		Data: str.Data,
		Len:  str.Len,
		Cap:  str.Len,
	}
}
