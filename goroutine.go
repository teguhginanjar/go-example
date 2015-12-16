package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func(x int) {
		for i := x; i > 0; i-- {
			fmt.Println("go 1:", i)
			time.Sleep(time.Millisecond * 500)
		}
		close(ch)
	}(10)
	<-ch

	ch2 := make(chan bool)
	go func(x int) {
		for i := x; i > 0; i-- {
			fmt.Println("go 2:", i)
			time.Sleep(time.Millisecond * 400)
		}
		ch2 <- true
	}(4)
	<-ch2

	ch3, ch4 := make(chan bool), make(chan bool)
	go func(x int) {
		for i := x; i > 0; i-- {
			fmt.Println("go 3: ", i)
		}
		ch3 <- true
	}(5)
	go func(x int) {
		for i := x; i > 0; i-- {
			fmt.Println("go 4: ", i)
		}
		ch4 <- true

	}(5)

	<-ch3
	<-ch4
}
