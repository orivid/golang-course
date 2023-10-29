package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {

	wg.Add(2)
	go sayHello()
	go sayHi()
	fmt.Println("Starting to do something")
	wg.Wait()
	fmt.Println("This is the end")
	wg.Add(1)
	go sayHello()
	wg.Wait()
}

func sayHello() {
	time.Sleep(time.Duration(2) * time.Second)
	fmt.Println("Hellloooo")
	wg.Done()
}

func sayHi() {
	time.Sleep(time.Duration(3) * time.Second)
	fmt.Println("Hiiiii")
	wg.Done()
}
