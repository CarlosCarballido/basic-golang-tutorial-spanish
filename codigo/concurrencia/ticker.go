package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for i := 1; i <= 3; i++ {
		<-ticker.C
		fmt.Println("pulso", i)
	}
}
