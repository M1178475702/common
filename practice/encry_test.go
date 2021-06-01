package practice

import (
	"crypto/md5"
	"fmt"
	"io"
	"testing"
)

func Md5(raw string) (ret string, err error) {
	m := md5.New()
	_, err = io.WriteString(m, raw)
	if err != nil {
		return
	}
	ret = string(m.Sum(nil))
	return
}

func TestMd5(t *testing.T) {
	fmt.Println(Md5("123456"))
}
