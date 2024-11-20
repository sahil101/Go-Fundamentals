package main

import "fmt"

func sum[T int](s []T, c chan T) {
	var sum T

	for _, v := range s {
		sum = sum + v
	}
	c <- sum // sum send to c
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func fibonacci_2(c, quit chan int) {
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
	// channels
	// By default, sends and receives block until the other side is ready. This allows goroutines to synchronize without explicit locks or condition variables.

	ch := make(chan int)

	go func() {
		ch <- 353
	}()

	value := <-ch

	fmt.Println(value)

	s := []int{213, 213, 23, 323, 656}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c

	fmt.Println(x, y, x+y)

	// Buffered channels
	bch := make(chan int, 2)

	// value_bch, err := <-bch

	// receives throw error when buffer is empty
	// if err {
	// 	panic("Error Happens")
	// }
	// fmt.Println(value_bch)

	bch <- 100
	bch <- 300
	// bch <- 400 // should throw error when the buffer is full

	// Range and close

	bch_2 := make(chan int, 10)
	go fibonacci(cap(bch_2), bch_2)
	for val := range bch_2 {
		fmt.Println(val)
	}

	// Select
	/* The select statement lets a goroutine wait on multiple communication operations.

	A select blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready.
	*/

	c_2 := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c_2)
		}
		quit <- 0
	}()
	fibonacci_2(c_2, quit)

}
