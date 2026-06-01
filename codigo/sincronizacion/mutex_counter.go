package main

import (
	"fmt"
	"sync"
)

type Contador struct {
	mu sync.Mutex
	n  int
}

func (c *Contador) Inc() {
	c.mu.Lock()
	c.n++
	c.mu.Unlock()
}

func (c *Contador) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.n
}

func main() {
	var c Contador
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Inc()
		}()
	}
	wg.Wait()
	fmt.Println(c.Value())
}
