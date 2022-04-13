package main

import (
	"fmt"
	"sync"
)

type ConMap struct {
	sync.RWMutex
	m map[interface{}]interface{}
}

//метод получения по ключу
func (m *ConMap) Get(key interface{}) (interface{}, bool) {
	//блокировка чтения
	m.RLock()
	v, ok := m.m[key]
	//разблокировка чтения
	m.RUnlock()
	return v, ok
}

func (m *ConMap) Set(key interface{}, value interface{}) {
	//блокировка чтения И записи
	m.Lock()
	m.m[key] = value
	//разблокировка чтения И записи
	m.Unlock()
}
func main() {
	arr1 := []string{"poopin'", "kweelin'", "yahoin'", "selebraitin'", "snifin' around", "smilin'"}
	arr2 := []string{"yahoin'", "selebraitin'"}

	wg := &sync.WaitGroup{}
	wg.Add(2)
	m := &ConMap{m: make(map[interface{}]interface{})}
	go func() {
		for _, e := range arr1 {
			_, ok := m.Get(e)
			if !ok {
				m.Set(e, struct{}{})
			} else {
				fmt.Println(e)
			}
		}
		wg.Done()
	}()
	go func() {
		for _, e := range arr2 {
			_, ok := m.Get(e)
			if !ok {
				m.Set(e, struct{}{})
			} else {
				fmt.Println(e)
			}
		}
		wg.Done()
	}()
	wg.Wait()
}
