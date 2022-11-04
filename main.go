package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	x := 1
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for i := 1; i < 101; i++ {
			ch <- x
			if i%2 == 1 {
				fmt.Println(i)
			}
		}
	}()
	go func() {
		for i := 1; i < 101; i++ {
			<-ch
			if i%2 == 0 {
				fmt.Println(i)
			}
		}
	}()
	wg.Wait()
}
