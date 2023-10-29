package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan string)

	go func() {
		time.Sleep(time.Duration(1) * time.Second)
		c <- "ori"
	}()

	fmt.Println("Continuing")
	result, ok := <-c

	close(c)

	if res, ok := <-c; !ok {
		fmt.Println("The channel is closed and the value received is", res)
	}

	fmt.Println(result, ok)
}
