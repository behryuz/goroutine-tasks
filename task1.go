package main

import (
	"fmt"
	"time"
)

// 1. Напиши функцию printSlowly(name string): печатает "Начали: <name>", ждёт time.Sleep(1 * time.Second),
// печатает "Готово: <name>". Вызови её три раза подряд обычным способом (без go) и выведи,
// сколько всего заняло времени (time.Now() / time.Since()).

func printSlowly(name string) {
	fmt.Println("Начали:", name)
	time.Sleep(1 * time.Second)
	fmt.Println("Готово:", name)
}

func task1() {
	t := time.Now()
	printSlowly("First")
	printSlowly("Second")
	printSlowly("Third")

	fmt.Println(time.Since(t))
}
