package sub

import (
	"common/cgo/util"
	"fmt"
)

/*
将有关C的类型的函数封装在一个包下，这样可以避免类型转化
此外，该包也可以定义同一个
 */

func InvokeQsort(){

	arr := []int32{2,3,4}
	cip := util.SliceToCIntPtr(arr)
	rp := util.Qsort(cip, len(arr))
	s := util.CIntPtrToSlice(rp, len(arr), cap(arr))
	fmt.Println(s)
}
