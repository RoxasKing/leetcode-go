package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go fibonacci(c, quit)
	i := 0
	for elem := range c {
		fmt.Println(elem)
		i++
		if i > 13 {
			close(quit)
			break
		}
	}
}
