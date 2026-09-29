package main

import (
	"fmt"
	"sync"
	"time"
)

// 5. Напиши функцию greet(name string, wg *sync.WaitGroup),
// которая через defer wg.Done() печатает "Привет, <name>!"
// после time.Sleep(300 * time.Millisecond). В main заведи срез из 5 имён
// и в цикле запусти горутину для каждого имени (не забудь Add и Wait).

func task5() {
	t := time.Now()

	wg := sync.WaitGroup{}
	names := []string{"Alice", "Bob", "Charlie", "David", "Eve"}
	for _, name := range names {
		wg.Add(1)
		greet(name, &wg)
	}

	wg.Wait()

	fmt.Println(time.Since(t))
}

func greet(name string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(300 * time.Millisecond)
	fmt.Println("Привет,", name)
}
