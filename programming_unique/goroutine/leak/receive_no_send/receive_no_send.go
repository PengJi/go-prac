package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	defer func() {
		time.Sleep(time.Second)
		fmt.Println("the number of gorountines: ", runtime.NumGoroutine())
	}()

	var ch chan struct{}
	go func() {
		ch <- struct{}{}
	}()
}
