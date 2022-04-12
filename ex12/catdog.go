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
	fmt.Print("Repeats: ")
	for v := range ch {
		if v != nil {
			fmt.Print(" ", v)
		}
	}
	fmt.Print("\n")
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
	m    map[interface{}]struct{}
}

//такой же конвейер, соответствие условию множества происходит за счет использования мапы
// которая может хранить только уникальные ключи
func (m *Model) decider(str interface{}) interface{} {
	_, ok := m.m[str]
	m.m[str] = struct{}{}
	if ok {
		return str
	}
	return nil
}

func main() {
	//t := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	m := Model{
		data: []interface{}{},
	}
	m.f = m.decider
	m.m = make(map[interface{}]struct{})
	m.data = append(m.data, words...)
	Convrun(&m)
	fmt.Print("Uniq: ")
	for k := range m.m {
		fmt.Print(" ", k)
	}
	fmt.Print("\n")
}

//https://www.random.org/strings/?num=100&len=2&loweralpha=on&unique=off&format=html&rnd=new
