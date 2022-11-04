package main

import "fmt"

func main() {
	var intChan chan int64
	var x int64
	intChan = make(chan int64, 50000)
	intChan <- 1
	for i := 0; i < 50000; i++ {
		<-intChan
		x++
		intChan <- x
	}
	for i := 0; i < 50000; i++ {
		<-intChan
		x++
		intChan <- x
	}
	x = <-intChan
	fmt.Println(x)
}
