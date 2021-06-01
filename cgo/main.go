package main


import "C"
import (
	"common/cgo/util"
	"fmt"
)

func main() {

	//arr := []int32{3,4,5,6,7}
	//arrp := util.SliceToCIntPtr(arr)
	//cintp := C.qsort(arrp, C.int(len(arr)))
	//s := util.CIntPtrToSlice(cintp, len(arr), cap(arr))
	////打印：c的print慢于go之后的println
	//fmt.Println(s)

	//sub.InvokeQsort()
	//util.SayHello()

	fmt.Println(util.BytesToString([]byte("1234")))
	var m map[int]int
	if _, ok := m[1]; ok {
		//就是map是nil也不会报错
	}
}
