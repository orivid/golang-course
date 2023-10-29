package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

var mut sync.Mutex

var counter int64

func main() {
	incrementAtomic()
}

func incrementMutex() {

	wg.Add(100)
	for i := 0; i < 100; i++ {
		mut.Lock()
		num := counter
		num++
		counter = num
		mut.Unlock()
		wg.Done()
	}
	wg.Wait()
	fmt.Println("The counter is", counter)
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}

func incrementAtomic() {
	wg.Add(100)
	for i := 0; i < 100; i++ {
		atomic.AddInt64(&counter, 1)
		wg.Done()
	}
	wg.Wait()
	fmt.Println("The counter is", counter)
}
