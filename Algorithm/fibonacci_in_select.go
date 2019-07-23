package Algorithm

import "fmt"

func Fibonacci(n int) {
	c := make(chan int)
	quit := make(chan int)
	go func() {
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
	}()
	i := 0
	for elem := range c {
		fmt.Println(elem)
		i++
		if i > n {
			close(quit)
			break
		}
	}
}
