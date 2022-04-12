package main

import (
	"fmt"
	"sync"
)

type ConMap struct {
	sync.RWMutex
	m map[interface{}]interface{}
}

// весь вывод в программе нужен для наблюдения за очередностью
// все слипы в программе нужны для исксутвенной очередности запуска горутин

//метод получения по ключу
func (m *ConMap) Get(key interface{}) (interface{}, bool) {
	//блокировка чтения
	m.RLock()
	v, ok := m.m[key]
	//разблокировка чтения
	m.RUnlock()
	return v, ok
}

func (m *ConMap) Set(key, value interface{}) {
	//блокировка чтения И записи
	m.Lock()
	m.m[key] = value
	//разблокировка чтения И записи
	m.Unlock()
}

type Conv interface {
	Pusher(ch chan interface{}, wg *sync.WaitGroup)
	Puller(ch chan interface{}, wg *sync.WaitGroup)
	Worker(ch_pull, ch_push chan interface{}, wg *sync.WaitGroup)
}

func (model *Model) Worker(ch_pull, ch_push chan interface{}, wg *sync.WaitGroup) {
	for v := range ch_pull {
		ch_push <- model.f(v)
	}
	close(ch_push)
	wg.Done()
}

//[]int{1, 2, 3, 4, 5, 6, 11111111111111, 645}
//пишет данные в канал и знает когда закончатся значения после чего закрывает канал
func (model *Model) Pusher(ch chan interface{}, wg *sync.WaitGroup) {
	x := model.data
	for _, e := range x {
		ch <- e
	}
	close(ch)
	wg.Done()
}

//послушно печатает ни о чем не задумываясб
func (model *Model) Puller(ch chan interface{}, wg *sync.WaitGroup) {
	for v := range ch {
		if v != nil {
			fmt.Println(v)
		}
	}
	wg.Done()
}

func Convrun(conv Conv) {
	var wg sync.WaitGroup
	c1 := make(chan interface{})
	c2 := make(chan interface{})
	wg.Add(3)
	go conv.Pusher(c1, &wg)
	go conv.Worker(c1, c2, &wg)
	go conv.Puller(c2, &wg)
	wg.Wait()
}

type Model struct {
	data []interface{}
	f    func(interface{}) interface{}
	m    *ConMap
}

//такой же конвейер только в случае повторного вхождения в мапу выбрасывает значение
//на экран, а можно кидать в массив
func (m *Model) decider(str interface{}) interface{} {
	i, ok := m.m.Get(str)
	if !ok {
		m.m.Set(str, 1)
	} else {
		m.m.Set(str, i.(int)+1)
		return str
	}
	return nil
}

func main() {
	//t := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	m1 := Model{
		data: []interface{}{},
		m:    &ConMap{m: make(map[interface{}]interface{})},
	}
	m1.data = append(m1.data, "poopin'", "kweelin'", "yahoin'", "selebraitin'", "snifin' around", "smilin'")
	m1.f = m1.decider
	m2 := Model{
		data: []interface{}{},
		m:    m1.m,
	}
	m2.data = append(m2.data, "yahoin'", "selebraitin'")
	m2.f = m2.decider
	Convrun(&m1)
	Convrun(&m2)
	fmt.Println(m1.m)
}
