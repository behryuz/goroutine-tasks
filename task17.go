package main

import (
	"fmt"
	"sync"
)

// 	17. Дан срез из 6 сайтов (строк). Заведи results := make([]string, len(sites)).
//	Запусти горутину на каждый сайт: горутина номер i после небольшой паузы
//	записывает results[i] = sites[i] + " - OK".
//	Дождись всех через WaitGroup и в конце выведи results по порядку.

func task17() {
	sites := []string{"google.com", "google2.com", "google3.com", "google4.com", "google5.com", "facebook.com"}
	results := make([]string, len(sites))
	var wg sync.WaitGroup

	for i := range sites {
		wg.Go(func() {
			results[i] = sites[i] + " - OK"
		})
	}

	wg.Wait()

	fmt.Println("results:", results)
}
