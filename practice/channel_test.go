package practice

import (
	"fmt"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	ch := make(chan int)
	close(ch)
	//close后，可以重复读，会返回零值，以及ok指示
	tmp, ok := <- ch
	fmt.Println(tmp, ok)
	ch <- 1
	tmp, ok = <- ch
	fmt.Println(tmp, ok)
	var chnil chan int
	go func() {
		fmt.Println("1")
		chnil <- 1
	}()
	fmt.Println(tmp, ok)
	time.Sleep(1*time.Second)
	tmp, ok = <- chnil
	fmt.Println(tmp, ok)
}


