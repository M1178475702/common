package practice

import (
	"fmt"
	"sync"
	"testing"
)

func TestCrossPrint(t *testing.T) {

	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		for {
			count := <- ch
			fmt.Println(count)
			if count >= 10 {
				wg.Done()
			}
			ch <- count+1
		}

	}()

	go func() {
		for {
			count := <- ch
			fmt.Println(count)
			if count >= 10 {
				wg.Done()
			}
			ch <- count+1
		}


	}()
	ch <- 1
	wg.Wait()

}

func TestBuf(t *testing.T) {
	bufch := make(chan struct{}, 2)

	go func() {
		bufch <- struct{}{}
	}()

	go func() {
		<- bufch
	}()

}

func TestSync1(t *testing.T) {

	one22ch := make(chan int, 1)
	one23ch := make(chan int, 1)
	two21ch := make(chan int, 1)
	thr21ch := make(chan int, 1)
	endch := make(chan struct{})
	//1
	go func() {
		cnt := 0
		for {
			fmt.Printf("prod %v\n", cnt)
			one22ch <- cnt
			one23ch <- cnt

			<-two21ch
			<-thr21ch
			cnt++
			if cnt == 10 {
				endch <- struct{}{}
				return
			}

		}
	}()

	//2
	go func() {
		for{
			cnt := <- one22ch
			fmt.Printf("c2 %v\n", cnt)
			two21ch <- 1
		}
	}()

	//3
	go func() {
		for{
			cnt := <- one23ch
			fmt.Printf("c3 %v\n", cnt)
			thr21ch <- 1
		}
	}()

	<-endch
}
