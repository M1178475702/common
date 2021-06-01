package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	go func() {
		ch := make(chan os.Signal)
		signal.Notify(ch, os.Interrupt)
		//注册了INT事件，便不会触发默认动作
		s := <-ch
		fmt.Printf("recv %v\n", s)
	}()
	//sync.Mutex{}
	for {
		;
	}
}
