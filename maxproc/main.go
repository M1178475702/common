package main

import (
	"fmt"
	"runtime"
	"sync"
)

//GOMAXPROCS
func main(){
	runtime.GOMAXPROCS(2)
	wg := sync.WaitGroup{}
	wg.Add(3)
	go func() {
		i := 1
		for i < 30 {
			fmt.Println("a", i)
			i++
		}
		wg.Done()
	}()
	go func() {
		i := 1
		for i < 30 {
			fmt.Println("b", i)
			i++
		}
		wg.Done()
	}()

	go func() {
		i := 1
		for i < 30 {
			fmt.Println("c", i)
			i++
		}
		wg.Done()
	}()
	wg.Wait()
}
