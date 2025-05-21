package main

import (
	"fmt"
	"time"
)


func generateNumbers(size int, out chan<- int) {
	for i := 1; i <= size; i++ {
		out <- i
	}
	close(out) 
}

func filter(in <- chan int, out chan<- int) {
	for i := range in {
		if i % 2 == 0  {
			out <- i;
		}
	}
	close(out)
}

func sum(in <- chan int, out chan<- int) {
	sum := 0;
	for i := range in {
		sum += i
	}

	out <- sum
	close(out)
}

func pipeline(size int) int {
	gen := make(chan int, 10)
	fil := make(chan int, 10)
	res := make(chan int)
	
	go generateNumbers(size, gen)
	go filter(gen, fil)
	go sum(fil, res)

	return <-res
}

func sequenece(size int) int {
	sum := 0
	for i := 1; i <= size; i++ {
		if i % 2 == 0 {
			sum += i;
		}
	}
	return sum;
}

func main() {
	start := time.Now() 
	res1 := pipeline(100_000)
	duration := time.Since(start)

	fmt.Println("Pipeline")
	fmt.Println("Res:", res1)
	fmt.Println("Duration:", duration)

	start = time.Now()
	res2 := sequenece(100_000)
	duration = time.Since(start)

	fmt.Println("Sequence")
	fmt.Println("Res:", res2)
	fmt.Println("Duration:", duration)
}