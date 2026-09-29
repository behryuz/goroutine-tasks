package main

import (
	"fmt"
	"sync"
)

// 10. Напиши программу с 10 горутинами, каждая печатает свой номер от 1 до 10 (fmt.Print).
// Дождись через WaitGroup.
// Запусти программу 3 раза подряд и сравни вывод — порядок чисел должен быть разным.

func task10() {
	wg := sync.WaitGroup{}

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go printNumber(i, &wg)
	}

	wg.Wait()
}

func printNumber(n int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(n)
}
