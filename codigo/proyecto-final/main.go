package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Task struct {
	ID int
}

type Result struct {
	TaskID int
	Value  int
}

type Stats struct {
	mu        sync.Mutex
	Completed int
}

func (s *Stats) IncCompleted() {
	s.mu.Lock()
	s.Completed++
	s.mu.Unlock()
}

func worker(ctx context.Context, tasks <-chan Task, results chan<- Result, stats *Stats, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case task, ok := <-tasks:
			if !ok {
				return
			}
			time.Sleep(100 * time.Millisecond)
			results <- Result{TaskID: task.ID, Value: task.ID * task.ID}
			stats.IncCompleted()
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	tasks := make(chan Task)
	results := make(chan Result)
	var stats Stats
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go worker(ctx, tasks, results, &stats, &wg)
	}

	go func() {
		defer close(tasks)
		for i := 1; i <= 8; i++ {
			select {
			case <-ctx.Done():
				return
			case tasks <- Task{ID: i}:
			}
		}
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Printf("task %d -> %d\n", result.TaskID, result.Value)
	}

	fmt.Printf("completed=%d\n", stats.Completed)
}
