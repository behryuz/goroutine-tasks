package main

import (
	"fmt"
	"sync"
	"time"
)

// 7. Перепиши задание 6 без отдельных функций — используй анонимную горутину прямо в main:
// go func(n int) { defer wg.Done(); ... }(5) и так же для 3 и 7.

func task7() {
	t := time.Now()

	wg := sync.WaitGroup{}

	count := func(n int, wg *sync.WaitGroup) {
		defer wg.Done()
		for i := range n {
			fmt.Printf("%d ", i+1)
			time.Sleep(10 * time.Millisecond)
		}

		fmt.Println()
	}

	wg.Add(1)
	count(5, &wg)
	wg.Add(1)
	count(3, &wg)
	wg.Add(1)
	count(7, &wg)

	wg.Wait()

	fmt.Println(time.Since(t))
}
