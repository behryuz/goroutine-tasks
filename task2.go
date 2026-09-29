package main

import (
	"fmt"
	"time"
)

// 2. Возьми функцию из задания 1 и запусти те же три вызова через go printSlowly(...),
// больше ничего не меняя. Запусти программу несколько раз.
// Комментарием в коде коротко напиши, что видишь в выводе (печатается ли вообще что-то и почему).

func task2() {
	t := time.Now()
	go printSlowly("First")
	go printSlowly("Second")
	go printSlowly("Third")

	fmt.Println(time.Since(t))
}
