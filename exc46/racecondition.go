package main

import (
	"fmt"
	"runtime"
	"sync"
)

var num int
var wg sync.WaitGroup
var mutex sync.Mutex

func main() {

	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.Compiler)
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go increment()
	}

	wg.Wait()
	fmt.Println("The num is", num)
}

func increment() {
	mutex.Lock()
	num++
	wg.Done()
	mutex.Unlock()
}
