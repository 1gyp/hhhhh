package main

import (
	"fmt"
	"sync"
)

var x int64
var wg sync.WaitGroup
var mu sync.Mutex

func add(ch chan int64) {
	var i int64
	for i = 0; i < 50000; i++ {
	}
	ch <- i
	wg.Done()
}
func main() {
	wg.Add(2)
	channel := make(chan int64)
	go add(channel)
	go add(channel)
	x := <-channel + <-channel
	wg.Wait()
	fmt.Println(x)
}
