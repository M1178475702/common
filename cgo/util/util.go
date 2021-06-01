package util
/*
#include <stdio.h>

void SayHello();
int* gsort(int* arr, int len);
void helloString(_GoString_ p0);
//void helloSlice(GoSlice p0);
*/
import "C"
import (
	"fmt"
	"reflect"
	"unsafe"
)


//导入函数后，会生成_cgo_export.h文件，该文件又会引入stdlib.h，从而导致qsort的重复声明

//export helloString
func helloString(s string) {
	fmt.Println(s)
}

//export add
func add(a, b C.int) C.int {
	return a + b
}

////export helloSlice
//func helloSlice(s []byte) {}

func SayHello(){
	C.SayHello()
	C.helloString("ddd")
}

//todo 引用的c文件要在同一包目录下

//64位平台下默认Int为64位，而c中int为32位
//传入[]int的内存分布为  v1 0 v2 0 v3 0
//slice ptr -> pointer -> slice_header_ptr
//unsafe.Pointer可以类型转化，uintptr可以运算，但不可以转化
func SliceToCIntPtr(s []int32) *C.int {
	sp := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	return (*C.int)(unsafe.Pointer(sp.Data))
}

//该函数还无法被其他包使用：*C.int => *main._Ctype_int
func CIntPtrToSlice(cip *C.int, length, capacity int) []int32 {
	//构造一个slice  *SliceHeader == &Slice
	sh := &reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(cip)),
		Len:  length,
		Cap:  capacity,
	}
	//类型转化只能用unsafe ptr转化
	//先将地址转化成pointer，再转化成其他类型指针（pointer做类型转化可以绕过类型检查）
	return *(*[]int32)(unsafe.Pointer(sh))
}


func BytesToString(bs []byte) string {
	//sh := reflect.StringHeader{
	//	Data: uintptr(unsafe.Pointer(&bs)),
	//	Len:  len(bs),
	//}
	//sh ptr => string ptr => evaluate
	//return *(*string)(unsafe.Pointer(&sh))
	//bs ptr => pointer => *string => evaluate
	return *(*string)(unsafe.Pointer(&bs))
}

func StringToBytes(s string) []byte {
	strhp := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: strhp.Data,
		Len:  strhp.Len,
		Cap:  strhp.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}


func Qsort(arr *C.int, length int) *C.int {
	return C.gsort(arr, C.int(length))
}
