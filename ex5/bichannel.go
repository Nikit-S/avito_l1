package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

//писатьель — функция для записи в канал
func writer(mch chan interface{}, quit chan bool, wg *sync.WaitGroup) {
	wg.Add(1)
	for {
		time.Sleep(50 * time.Millisecond)
		select {
		//постоянно пишет в основной канал если не выполняется кейс
		default:
			rand.Seed(time.Now().UnixNano())
			//подача рандомного символа в канал
			mch <- byte('A' + rand.Intn(10))
		// выполняется при закрытии канла или при подачи значения в канал
		case <-quit:
			wg.Done()
			return
		}
	}
}

//читатель — функция для чтения из канала
func reader(mch chan interface{}, quit chan bool, wg *sync.WaitGroup) {
	wg.Add(1)
	for {
		//задержка для корректной работы терминала
		time.Sleep(50 * time.Millisecond)
		select {
		//либо горутина получит данные в канал дата, либо будедт закрытие канала quit
		case data := <-mch:
			fmt.Printf("%c", data)
		case <-quit:
			wg.Done()
			return
		}
	}
}

func main() {
	var n int
	var wg sync.WaitGroup
	fmt.Println("How long do I suppose to wait?")
	fmt.Scan(&n)

	if n <= 0 {
		return
	}
	mch := make(chan interface{}, 1)
	quit := make(chan bool, 1)

	go writer(mch, quit, &wg)
	go reader(mch, quit, &wg)

	log.Println("waitin'")
	time.Sleep(time.Duration(n) * time.Second)
	//используется закрытие чтобы не кидать вкаждый канал значение
	close(quit)
	wg.Wait()
	fmt.Println("")

	//на случай незкарытого канала и данных в нем проверка на открытость а вообще

	/*You needn't close every channel when you've finished with it.
	It's only necessary to close a channel when it is important to tell the
	receiving goroutines that all data have been sent. A channel that the
	garbage collector determinies to be unreachable will have
	its resources reclaimed whether or not it is closed.
	*/
	select {
	case _, ok := <-mch:
		if ok {
			log.Println("closin'")
			close(mch)
		}
	default:
	}
	log.Println("exit'")

}
