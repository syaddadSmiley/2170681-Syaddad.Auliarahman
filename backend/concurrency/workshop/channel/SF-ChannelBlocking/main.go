package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan string)
	go func() {
		time.Sleep(time.Second * 4)
		c1 <- 1
	}()
	go func() {
		time.Sleep(time.Second * 5)
		c2 <- "two"
	}()
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
			// default:
			// 	fmt.Println("Default")
		}
	}
}
