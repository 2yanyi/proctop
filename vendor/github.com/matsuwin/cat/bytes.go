package cat

import (
	"reflect"
	"unsafe"
)

/**
 * Provides low-level type conversion without generating a memory copy.
 *
 * 该程序代码提供底层类型转换，不会产生内存拷贝。
 *
 */

// BytesToString is equivalent to string(a)
func BytesToString(a []byte) string {
	if len(a) > _size {
		return *(*string)(unsafe.Pointer(&a))
	}
	return string(a)
}

// RunesToString is equivalent to string(a)
func RunesToString(a []rune) string {
	if len(a) > _size {
		return *(*string)(unsafe.Pointer(&a))
	}
	return string(a)
}

// StringToBytes is equivalent to []byte(a)
func StringToBytes(a *string) []byte {
	if len(*a) > _size {
		return *(*[]byte)(unsafe.Pointer(str2slice((*reflect.StringHeader)(unsafe.Pointer(a)))))
	}
	return []byte(*a)
}

// StringToRunes is equivalent to []rune(a)
func StringToRunes(a *string) []rune {
	if len(*a) > _size {
		return *(*[]rune)(unsafe.Pointer(str2slice((*reflect.StringHeader)(unsafe.Pointer(a)))))
	}
	return []rune(*a)
}

func str2slice(str *reflect.StringHeader) *reflect.SliceHeader {
	return &reflect.SliceHeader{
		Data: str.Data,
		Len:  str.Len,
		Cap:  str.Len,
	}
}

const _size = 32
