package main

import (
	"fmt"
	"sync"
	"time"
)

// 	16. Дан срез из 5 имён файлов. Функция download(name string, wg *sync.WaitGroup)
//	печатает "Скачали: <name>" после паузы 1 секунда.
//	Запусти скачивание всех файлов одновременно через горутины и WaitGroup.
// 	Выведи время работы и сравни с тем, сколько заняло бы скачать файлы по одному (5 секунд).

func task16() {
	files := []string{"file1.txt", "file2.txt", "file3.txt", "file4.txt", "file5.txt"}
	t := time.Now()
	var wg sync.WaitGroup

	for _, f := range files {
		wg.Add(1)
		go download(f, &wg)
	}

	wg.Wait()

	fmt.Println("time:", time.Since(t))
}

func download(name string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(1 * time.Second)
	fmt.Println("Downloaded", name)
}
