package main


import "C"
import (
	"common/cgo/sub"
	"common/cgo/util"
)

func main() {

	//arr := []int32{3,4,5,6,7}
	//arrp := util.SliceToCIntPtr(arr)
	//cintp := C.qsort(arrp, C.int(len(arr)))
	//s := util.CIntPtrToSlice(cintp, len(arr), cap(arr))
	////打印：c的print慢于go之后的println
	//fmt.Println(s)

	sub.InvokeQsort()
	util.SayHello()
}
