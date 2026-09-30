package main

import (
	"fmt"
	"sync"
)

// 	18. Дан срез из 10 чисел. Заведи doubled := make([]int, 10).
//	Запусти горутину на каждое число: она умножает своё число на 2
//	и кладёт результат в doubled по тому же индексу (каждая горутина трогает только свою ячейку).
//	После wg.Wait() выведи рядом исходный срез и doubled.

func task18() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	doubled := make([]int, 10)

	var wg sync.WaitGroup

	for i := range nums {
		wg.Go(func() {
			doubled[i] = nums[i] * 2
		})
	}

	wg.Wait()

	fmt.Println("nums:", nums)
	fmt.Println("doubled:", doubled)
}
