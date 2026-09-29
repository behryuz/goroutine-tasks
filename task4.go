package main

import (
	"fmt"
	"sync"
	"time"
)

// 4. Перепиши задание 3 без time.Sleep в main — с помощью sync.WaitGroup:
// заведи var wg sync.WaitGroup; перед каждым go printSlowly(...)
// вызови wg.Add(1); передай &wg в функцию; внутри функции первой
// строкой поставь defer wg.Done(); в конце main вызови wg.Wait().

func task4() {
	t := time.Now()

	wg := sync.WaitGroup{}

	wg.Add(1)
	go printSlowlyWg("First", &wg)
	wg.Add(1)
	go printSlowlyWg("Second", &wg)
	wg.Add(1)
	go printSlowlyWg("Third", &wg)

	wg.Wait()

	fmt.Println(time.Since(t))
}

func printSlowlyWg(name string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Начали:", name)
	fmt.Println("Готово:", name)
}
