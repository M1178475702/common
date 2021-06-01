package practice

import (
	"fmt"
	"github.com/cockroachdb/errors"
	"testing"
)

func cleanPtr(err *error) {
	fmt.Println(*err)
}

func TestDeferWithPtr(t *testing.T) {

	var err error
	defer cleanPtr(&err)
	err = errors.New("ddd")
	fmt.Println(err)
}


func RetWithErr() (err error) {
	clean := func() {
		//使用闭包，而不是参数传值；参数传值会在defer调用时确定值
		//使用闭包便无法进行封装式处理，只能现场声明闭包函数，或用闭包将待传值函数进行包装
		if err != nil {

		}
	}
	defer clean()

	defer func() {
		cleanPtr(&err)
	}()
	return
}



