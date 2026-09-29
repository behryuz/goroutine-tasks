package main

import (
	"fmt"
	"sync"
)

// 13. Заведи переменную counter := 0 и запусти 1000 горутин,
// каждая делает counter++ (без всякой защиты, только WaitGroup).
// Выведи итоговое значение counter — запусти программу несколько раз и убедись,
// что число каждый раз разное и почти никогда не равно 1000.

func task13() {
	// ==================
	//WARNING: DATA RACE
	//Read at 0x00c00011c038 by goroutine 10:
	//  main.task13.func1()
	var counter int
	var wg sync.WaitGroup

	wg.Add(1000)

	for range 1000 {
		go func() {
			defer wg.Done()
			counter++
		}()
	}

	wg.Wait()

	fmt.Println("counter =", counter)
}
