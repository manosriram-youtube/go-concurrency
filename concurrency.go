package main

import (
	"fmt"
	"sync"
)

/*
	Mutex
		- RLock() -> multiple reads, single write
		- Lock() -> single read / write
	WaitGroups
*/

var waitg sync.WaitGroup
var Mutex = sync.RWMutex{}
var Counter = 0

func Hello() {
	fmt.Printf("Hello! Counter = %v\n", Counter)
	Mutex.RUnlock()
	defer waitg.Done()
}

func Increment() {
	Counter++
	Mutex.Unlock()
	defer waitg.Done()
}

func main() {
	for i := 0; i < 10; i++ {
		waitg.Add(2)

		Mutex.RLock()
		go Hello()

		Mutex.Lock()
		go Increment()
	}
	waitg.Wait()
}
