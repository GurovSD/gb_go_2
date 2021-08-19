package main

import "fmt"

func main() {

	ch_counter := make(chan int)
	var workers = make(chan struct{}, 1000)

	ch_counter <- 0

	for i := 0; i < 1000; i++ {
		workers <- struct{}{}

		go func() {
			defer func() {
				<-workers
			}()

			buf_val := <-ch_counter
			fmt.Println(buf_val)
			buf_val++
			ch_counter <- buf_val

		}()

	}

	fmt.Println(len(workers))

	for {
		if len(workers) == 0 {
			fin_value := <-ch_counter
			fmt.Println(fin_value)
			break
		}
	}

}
