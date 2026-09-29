package main

import (
	"fmt"
	"sync"
)

// 15. Возьми задание 13, но вместо одной общей переменной counter заведи results := make([]int, 1000).
// Пусть горутина номер i делает results[i] = 1 (только свою ячейку, не общую переменную).
// После wg.Wait() посчитай сумму среза обычным циклом.
// Убедись, что теперь результат всегда ровно 1000.

func task15() {
	results := make([]int, 1000)
	var wg sync.WaitGroup

	wg.Add(1000)

	for i := range 1000 {
		go func() {
			defer wg.Done()
			results[i] = 1
		}()
	}

	wg.Wait()

	sum := 0

	for i := range results {
		sum += results[i]
	}

	fmt.Println("sum:", sum)
}
