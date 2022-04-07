package main

import (
	"fmt"
	"sync"
)

type ConMap struct {
	sync.RWMutex
	m map[string]int
}

// весь вывод в программе нужен для наблюдения за очередностью
// все слипы в программе нужны для исксутвенной очередности запуска горутин

//метод получения по ключу
func (m *ConMap) Get(key string) (int, bool) {
	//блокировка чтения
	m.RLock()
	v, ok := m.m[key]
	//разблокировка чтения
	m.RUnlock()
	return v, ok
}

func (m *ConMap) Set(key string, value int) {
	//блокировка чтения И записи
	m.Lock()
	m.m[key] = value
	//разблокировка чтения И записи
	m.Unlock()
}

func main() {
	arr1 := []string{"poopin'", "kweelin'", "yahoin'", "selebraitin'", "snifin' around", "smilin'"}
	arr2 := []string{"yahoin'", "selebraitin'"}

	m := &ConMap{m: make(map[string]int)}
	var wg sync.WaitGroup
	wg.Add(2)

	//горутина записи
	go func(m *ConMap, wg *sync.WaitGroup, arr []string) {
		for _, e := range arr {
			i, ok := m.Get(e)
			if !ok {
				m.Set(e, 1)
			} else {
				m.Set(e, i+1)
			}
		}
		wg.Done()
	}(m, &wg, arr1)

	go func(m *ConMap, wg *sync.WaitGroup, arr []string) {
		for _, e := range arr {
			i, ok := m.Get(e)
			if !ok {
				m.Set(e, 1)
			} else {
				m.Set(e, i+1)
			}
		}
		wg.Done()
	}(m, &wg, arr2)

	wg.Wait()

	fmt.Println(m.m)

}
