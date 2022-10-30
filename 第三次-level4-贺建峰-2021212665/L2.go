package main

import (
	"fmt"
	"sync"
	"time"
)

var w sync.WaitGroup

func ad(ch chan int) {
	fmt.Printf("%d ", <-ch)
}
func main() {
	channel := make(chan int)
	for i := 1; i <= 100; i++ {
		go ad(channel)
		channel <- i //发送
		i++
		go ad(channel)
		channel <- i
		time.Sleep(time.Millisecond)
	}
}
