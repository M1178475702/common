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
	c.EXPECT().Get(gomock.Any()).AnyTimes().Return("return").Do(func() {
		fmt.Println("do")
	})
	fmt.Println(c.Get("key"))


}
