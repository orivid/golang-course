package main

import (
	"log"
	"reflect"
	"runtime"
	"time"
)

func main() {
	logOperation(4, doingSomething)
}

func logOperation(n int, f func(int)) {

	funcName := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	log.Printf("The method %v is starting", funcName)
	start := time.Now()
	f(n)
	sub := time.Since(start)
	log.Printf("The method %T has finished", funcName)
	log.Printf("It took the method %v to run", sub)

}

func doingSomething(n int) {
	time.Sleep(time.Duration(n) * time.Second)
}
