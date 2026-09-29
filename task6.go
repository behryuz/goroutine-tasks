package main

import (
	"fmt"
	"sync"
	"time"
)

// 6. Напиши функцию countTo(n int, wg *sync.WaitGroup), печатающую числа от 1 до n
// с паузой 10 * time.Millisecond между ними. Запусти как горутины countTo(5, ...),
// countTo(3, ...), countTo(7, ...) и дождись всех через WaitGroup.

func task6() {
	t := time.Now()

	wg := sync.WaitGroup{}
	wg.Add(1)
	countTo(5, &wg)
	wg.Add(1)
	countTo(3, &wg)
	wg.Add(1)
	countTo(7, &wg)

	wg.Wait()

	fmt.Println(time.Since(t))
}

func countTo(n int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := range n {
		fmt.Printf("%d ", i+1)
		time.Sleep(10 * time.Millisecond)
	}

	fmt.Println()
}
