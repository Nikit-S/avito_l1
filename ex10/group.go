package main

import (
	"fmt"
	"sync"
)

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
		fmt.Println(v)
	}
	wg.Done()
}

func Convrun(conv Conv) {
	fmt.Println("Ignite")
	var wg sync.WaitGroup
	c1 := make(chan interface{})
	c2 := make(chan interface{})
	wg.Add(3)
	fmt.Println("strting goroutines")
	go conv.Pusher(c1, &wg)
	go conv.Worker(c1, c2, &wg)
	go conv.Puller(c2, &wg)
	wg.Wait()
}

type Model struct {
	data []interface{}
	f    func(interface{}) interface{}
	m    map[int][]float64
}

/*
здесь использован конвейер из предыдущего номера с добавлением мапы и особой
функцией которая работает с мапой: а именно проверяет наличие необходимой группы
создает оную при необходимости после чего добавляет туда элемент на основе его
десятичного значения, для ускорения процесса например можно заупсутиь несколько
таких конвейеров но тогда нужно будет вешать замки
*/
func (m *Model) decider(f interface{}) interface{} {
	var v int
	if f.(float64) < 0 {
		v = int(f.(float64) - 10)
	} else {
		v = int(f.(float64))
	}

	group := v / 10 * 10
	fl := f.(float64)
	_, ok := m.m[group]
	if !ok {
		m.m[group] = make([]float64, 0)
	}
	m.m[group] = append(m.m[group], fl)
	return fl
}

func main() {
	//t := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	m := Model{
		data: make([]interface{}, 0),
		m:    make(map[int][]float64),
	}
	m.data = append(m.data, -25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, 1.0, -1.0, -9.1, 0.0)
	m.f = m.decider
	Convrun(&m)
	fmt.Println(m.m)
}
