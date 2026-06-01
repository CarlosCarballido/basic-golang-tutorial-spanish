package main

import "fmt"

func main() {
	jobs := make(chan int)
	results := make(chan int)

	go func() {
		defer close(jobs)
		for i := 1; i <= 4; i++ {
			jobs <- i
		}
	}()

	go func() {
		for job := range jobs {
			results <- job * job
		}
		close(results)
	}()

	for result := range results {
		fmt.Println(result)
	}
}
