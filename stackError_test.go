package common

import (
	"errors"
	"fmt"
	"testing"
)


func TestNewErr(t *testing.T) {
	src := errors.New("src err")
	err := NewErr(src, "second err")
	fmt.Println(WithStack(err))
	err = NewErr(err, "third err")
	fmt.Println(WithStack(err))
	fmt.Println(err)
}

func TestDefer(t *testing.T) {
	c := 1
	defer func() {
		fmt.Println(c)
		c += 1
	}()

	a := func() int{
		c += 1
		return c + 1
	}
	c += a()
}