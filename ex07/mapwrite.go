package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type ConMap struct {
	sync.RWMutex
	m map[int]string
}

// весь вывод в программе нужен для наблюдения за очередностью
// все слипы в программе нужны для исксутвенной очередности запуска горутин

//метод получения по ключу
func (m *ConMap) Get(key int) (string, bool) {
	log.Println("get lock of:", key)
	//блокировка чтения
	m.RLock()
	log.Println("get locked of:", key)
	time.Sleep(time.Duration(key*2) * time.Second)
	v, ok := m.m[key]
	//разблокировка чтения
	m.RUnlock()
	log.Println("Get UNlocked of:", key)
	return v, ok
}

func (m *ConMap) Set(key int, value string) {
	log.Println("Set lock of:", key)
	//блокировка чтения И записи
	m.Lock()
	log.Println("Set locked of:", key)
	time.Sleep(3 * time.Second)
	m.m[key] = value
	//разблокировка чтения И записи
	m.Unlock()
	log.Println("Set UNlocked of:", key)
}

//а вообще есть sync.Map который так и работает
func main() {
	m := &ConMap{m: make(map[int]string)}
	var wg sync.WaitGroup
	wg.Add(3)

	//горутина записи
	go func(m *ConMap, wg *sync.WaitGroup) {
		m.Set(1, "Hello")
		wg.Done()
	}(m, &wg)
	//горутина записи
	go func(m *ConMap, wg *sync.WaitGroup) {
		m.Set(1, "World")
		wg.Done()
	}(m, &wg)
	time.Sleep(1 * time.Second)

	//горутина чтения
	go func(m *ConMap, wg *sync.WaitGroup) {
		fmt.Println(m.Get(2))
		wg.Done()
	}(m, &wg)

	time.Sleep(1 * time.Second)

	//горутина чтения
	go func(m *ConMap, wg *sync.WaitGroup) {
		fmt.Println(m.Get(1))
		wg.Done()
	}(m, &wg)

	wg.Wait()

}
