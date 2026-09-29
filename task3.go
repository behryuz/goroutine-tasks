package main

import (
	"fmt"
	"time"
)

// 3. Почини задание 2 "костылём": в конце main добавь один time.Sleep(2 * time.Second),
// чтобы горутины успели закончиться.
// Комментарием напиши одно предложение — в чём минус такого способа ждать горутины.

func task3() {
	t := time.Now()
	go printSlowly("First")
	go printSlowly("Second")
	go printSlowly("Third")

	time.Sleep(2 * time.Second)

	fmt.Println(time.Since(t))
}
