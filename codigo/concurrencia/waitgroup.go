package main

import (
	"fmt"
	"sync"
)

func trabajo(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("trabajo", id)
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go trabajo(i, &wg)
	}
	wg.Wait()
}
