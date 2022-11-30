// Создайте две горутины, чтобы числа из одного канала читались по мере поступления, возводились в квадрат и результат записывался во второй канал.
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	origin := make(chan int)
	res := make(chan int)

	wg.Add(1)
	go func(in <-chan int, out chan<- int) {

		defer wg.Done()

		for r := range in {
			fmt.Printf("send result: %d\n", r*r)
			out <- r * r
		}
		close(out)

	}(origin, res)

	wg.Add(1)
	go func(out chan<- int) {

		defer wg.Done()

		for i := 0; i < 10; i++ {
			fmt.Printf("send number: %d\n", i)
			out <- i
		}
		close(out)

	}(origin)

	for r := range res {

		fmt.Printf("get result: %d\n", r)
	}

	wg.Wait()
}
