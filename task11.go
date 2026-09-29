package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// 11. Используй runtime.NumGoroutine(): напечатай его один раз до запуска 50 горутин
// (каждая просто спит 100 * time.Millisecond),
// один раз сразу после запуска всех (до wg.Wait()),
// и один раз после wg.Wait().

func task11() {
	fmt.Println("Перед запуском: ", runtime.NumGoroutine())
	wg := sync.WaitGroup{}
	for range 50 {
		wg.Add(1)
		go checkNumGoRoutine(&wg)
	}

	fmt.Println("Перед wait: ", runtime.NumGoroutine())

	wg.Wait()
	fmt.Println("После wait: ", runtime.NumGoroutine())
}

func checkNumGoRoutine(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(100 * time.Millisecond)
}
