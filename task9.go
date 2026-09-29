package main

import (
	"fmt"
	"sync"
	"time"
)

// 9. К заданию 8 добавь: перед циклом с горутинами сделай то же самое (8 пауз) обычным циклом, без go.
// Выведи оба времени — обычное и с горутинами — рядом друг с другом.

func task9() {
	t1 := time.Now()

	wg := sync.WaitGroup{}

	for i := 1; i <= 8; i++ {
		wg.Add(1)
		func(id int) {
			defer wg.Done()
			time.Sleep(time.Duration(id) * 50 * time.Millisecond)
			fmt.Printf("Обычный вызов %d завершена\n", id)
		}(i)
	}

	defaultTime := time.Since(t1)

	t2 := time.Now()
	for i := 1; i <= 8; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			time.Sleep(time.Duration(id) * 50 * time.Millisecond)
			fmt.Printf("Горутина %d завершена\n", id)
		}(i)
	}

	wg.Wait()

	goRoutineTime := time.Since(t2)

	fmt.Println("DefaultTime:", defaultTime)
	fmt.Println("GoRoutineTime:", goRoutineTime)
}
