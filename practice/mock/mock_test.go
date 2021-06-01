package mock

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"practive/test/imp"

	"testing"
)

func TestMock(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	c := imp.NewMockMock(ctrl)
	//c.EXPECT().Get(gomock.Any()).AnyTimes().Return("return").Do(func(key string) {
	//	fmt.Println("do", key)
	//})
	//
	//Do，签名要与函数本身相同
	c.EXPECT().Get(gomock.Any()).AnyTimes().DoAndReturn(func(key string) string {
		return "ret"
	})

	fmt.Println(c.Get("key"))
}
