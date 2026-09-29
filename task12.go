package main

import (
	"fmt"
	"sync"
	"time"
)

// 12. Запусти 100_000 горутин, каждая из которых делает только time.Sleep(1 * time.Millisecond).
// Дождись всех через WaitGroup и выведи, сколько заняло времени всего.

func task12() {
	t := time.Now()

	wg := sync.WaitGroup{}

	go func(n int, wg *sync.WaitGroup) {
		defer wg.Done()
		for range n {
			time.Sleep(1 * time.Millisecond)
		}
	}(100_000, &wg)

	wg.Wait()

	fmt.Println(time.Since(t))
}
