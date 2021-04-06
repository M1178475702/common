package common

import (
	"errors"
	"fmt"
	"runtime"
)

type StackError struct {
	err   error
	msg   string
	stack string
}

func (s *StackError) Error() string {	//原有项目直接使用err.Error作为错误信息返回前端，不适合处理
	return s.err.Error()
}

//可能改造方式：封装ajaxmsg
func (s *StackError) Stack() string {
	return fmt.Sprintf("%v", s.stack)
}

func (s *StackError) NestedErr() string {
	return ""
}

func WithStack(err error) string {
	//msg is used just in WithStack
	if s, ok := err.(*StackError); ok {
		return fmt.Sprintf("%v\n%v", s.msg, s.stack)
	} else {
		return fmt.Sprintf("%v", err)
	}
}

func NewErr(err error, msgf string, v ...interface{}) error {
	msg := fmt.Sprintf(msgf, v...)
	if err == nil {
		err = errors.New(msg)
		msg = ""
	}
	if src, ok := err.(*StackError); !ok {
		se := new(StackError)
		se.msg = msg
		buf := make([]byte, 512)
		se.err = err
		runtime.Stack(buf, false)
		se.stack = string(buf)
		return se
	} else {
		src.msg = fmt.Sprintf("%v\n%v", msg, src.msg)
		return src
	}
}


