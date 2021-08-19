package main

import "fmt"

func main() {

	ch_counter := make(chan int)
	ch_res := make(chan int)

	ch_counter <- 0

	go func() {

		val := <-ch_counter
		ch_res <- val

	}()

	val := <-ch_res
	fmt.Println(val)

}
