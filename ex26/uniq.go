package main

import (
	"fmt"
	"sync"
)

type ConMap struct {
	sync.RWMutex
	m map[rune]struct{}
}

// весь вывод в программе нужен для наблюдения за очередностью
// все слипы в программе нужны для исксутвенной очередности запуска горутин

//метод получения по ключу
func (m *ConMap) Get(key rune) (struct{}, bool) {
	//блокировка чтения
	m.RLock()
	v, ok := m.m[key]
	//разблокировка чтения
	m.RUnlock()
	return v, ok
}

func (m *ConMap) Set(key rune, value struct{}) {
	//блокировка чтения И записи
	m.Lock()
	m.m[key] = value
	//разблокировка чтения И записи
	m.Unlock()
}

func main() {

	word_source := []string{"lybjrEpTleTjYGcCjqzj", "nzeUAFoXWbnBLgcayFQO", "abcDeFghijklmnOpqrstuvwxyz"}
	for _, word := range word_source {
		m := &ConMap{m: make(map[rune]struct{})}

		//бежит слева направо
		var wg sync.WaitGroup
		wg.Add(2)
		go func() {
			for _, e := range word {
				_, ok := m.Get(e)
				//если элемент уже есть то выходит
				if ok {
					break
				} else {
					//приводит все к нижнему регистру
					if e > 'A' && e < 'Z' {
						e = e - 'A' + 'a'
					}
					m.Set(e, struct{}{})
				}
			}
			wg.Done()
		}()

		//бежит справа налево
		go func() {
			i := len(word) - 1
			for i >= 0 {
				e := rune(word[i])
				_, ok := m.Get(e)
				if ok {
					break
				} else {
					if e > 'A' && e < 'Z' {
						e = e - 'A' + 'a'
					}
					m.Set(e, struct{}{})
				}
				i--
			}
			wg.Done()
		}()
		wg.Wait()

		//если по итогу количество букв в слове совпадает с количеством уникальных
		//ключей в мапе =» повторов не было
		fmt.Println(len(word) == len(m.m))
	}

}
