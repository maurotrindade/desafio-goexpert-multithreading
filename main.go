package main

import (
	"fmt"
	"sync"
	"time"
)

func call(url string, wg *sync.WaitGroup, t int) {
	time.Sleep(time.Duration(t) * time.Millisecond)
	fmt.Print(url)
	wg.Done()
}

func main() {
	waitgroup := sync.WaitGroup{}
	waitgroup.Add(1)
	defer waitgroup.Wait()

	go call("primus.com", &waitgroup, 2)
	go call("secundus.com", &waitgroup, 3)
}
