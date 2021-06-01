package practice

import (
	se "errors"
	"fmt"
	"github.com/pkg/errors"
	"testing"
)

func retWithErr() error {
	return errors.Wrapf(se.New("nihao"), "hello")
}

func wrapErr(err error) error {

	return errors.Unwrap(err)
}

func TestError(t *testing.T) {

	fmt.Printf("%+v\n", retWithErr())  // hello: nihao
	fmt.Println()
	fmt.Printf("%+v\n", wrapErr(retWithErr()))
	//err1 stack1 format stack2
	//fmt.Println(errors.GetAllDetails(errors.WithStack(err)))  // []
	//de := errors.GetSafeDetails(err)
	//fmt.Println(de.SafeDetails) // [trace1, trace2, ...] （without err msg）
	//fmt.Println(errors.FlattenDetails(err))
}

/*
网络端连接报错后的处理：
1. 打日志，不处理
2. 重试，仍失败，打日志，不处理
3. 增加报警机制？
 */

/**
错误处理方案：
1. 是否每层都需要包装，或是实现定义一个底层，仅底层包装，其它层直接返回（或者仅with message）
- 每层都包装stack，stack会重复，无用
- with message的后端是： err1 stack msg，仍然很丑，不合适
- unwrap，去掉了堆栈信息
- 首先对项目分层，然后仅在最底层进行错误封装。此外，也包括三方包，std的错误。但这样需要人为识别该不该包装错误，项目小，分层好的还好说。
  若分层不好（或不好分层），以及项目过大，则区分是否应该包装错误会带来很大的负担。
- 对于暴露出去的三方包，不应该包装error。首先，外部使用者不应该关心代码是在三方包的哪里出错的。其次，会给使用者带来不确定的负担。
 */
