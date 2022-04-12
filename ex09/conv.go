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

//запуск конвейера
func Convrun(conv Conv) {
	var wg sync.WaitGroup
	//создание каналов записи и чтения
	c1 := make(chan interface{})
	c2 := make(chan interface{})
	wg.Add(3)
	//запуск трех рутин
	go conv.Pusher(c1, &wg)
	go conv.Worker(c1, c2, &wg)
	go conv.Puller(c2, &wg)
	wg.Wait()
}

//тут организован универсальный конвейер, который работает на основе данных
//предоставленных в Model, единсвтенным главным условием является наличие data
//и функции обработчика которая принимает в себя один элемент и отдает какой-то элемент
type Model struct {
	data []interface{}
	f    func(interface{}) interface{}
}

//to do chan 0

func main() {
	//t := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	m := Model{
		data: []interface{}{},
		f: func(i interface{}) interface{} {
			return i.(int) * 2
		},
	}
	m.data = append(m.data, 1, 2, 3, 4, 5, 6, 11111111111111, 645)
	Convrun(&m)
}
