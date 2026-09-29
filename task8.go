package main

import (
	"fmt"
	"sync"
	"time"
)

// 8. Заведи for i := 1; i <= 8; i++ и в цикле запускай анонимную горутину,
// которая получает номер как аргумент (go func(id int) {...}(i)),
// спит time.Duration(id) * 50 * time.Millisecond и печатает "Горутина <id> завершена".
// Дождись всех через WaitGroup.

func task8() {
	t := time.Now()

	wg := sync.WaitGroup{}

	for i := 1; i <= 8; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			time.Sleep(time.Duration(id) * 50 * time.Millisecond)
			fmt.Printf("Горутина %d завершена\n", id)
		}(i)
	}

	wg.Wait()

	fmt.Println(time.Since(t))
}
