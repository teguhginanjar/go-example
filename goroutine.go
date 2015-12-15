package main

import (
	"fmt"
	"time"
)

func main() {
	loop := make(chan int)

	go func(x int) {
		for i := x; i > 0; i-- {
			fmt.Println(i)
			time.Sleep(time.Millisecond * 500)
		}
		close(loop)
	}(10)
	<-loop
}
