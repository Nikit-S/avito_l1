package main

import (
	"fmt"
	"sync"
)

//реализация мьютекса внутри структуры которую
//впоследствии будут использовать несколько горутин
type Summ struct {
	//композиция типа
	int
	//композиция метода блокировки замка
	sync.Mutex
}

func main() {
	var wg sync.WaitGroup

	arr := []int{2, 4, 6, 8, 10}
	wg.Add(len(arr))

	var sum Summ
	for _, e := range arr {
		go func(e int, sum *Summ) {
			//все очень похоже на предыдущий номер
			// только здесь для избежания проблемы гонки
			// используется мьютекс, причем тот, тот который блокирует как
			// чтение так и запись
			sum.Lock()
			sum.int += e * e
			sum.Unlock()
			defer wg.Done()
		}(e, &sum)
	}
	wg.Wait()
	fmt.Println(sum.int)
}
