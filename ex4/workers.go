package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

//просто сложили все в одно место
type JobOffer struct {
	//основной канал «задач» для чтения
	mch chan interface{}
	//группа ожидания
	sync.WaitGroup
	//канал ожидания выхода
	quit chan bool
}

//функция инициализации переменных
func initwork(jo *JobOffer, n *int) bool {

	fmt.Print("Write number of workers: ")
	fmt.Scan(n)

	//проверка на натуральность
	if *n <= 0 {
		return false
	}

	//создание каналов
	jo.mch = make(chan interface{})
	jo.quit = make(chan bool, 1)
	return true
}

func work(jo *JobOffer, n int) {
	// каждый работник увеличивает группу ожидания на 1 (на одного себя)
	jo.Add(1)

	//инициализация переменной для получения в нее любого типа данных
	var task interface{}
	for {
		//тру рандом
		rand.Seed(time.Now().UnixNano())

		//селект позволяет выполнять участок кода в завоимости от канала который
		// получит значение первым
		select {
		// если повилась новая задача в списке задач
		case task = <-jo.mch:
			log.Println("Worker #", n, task)
			//делаем вид что у нас сложный процесс, который занимает время
			time.Sleep(time.Duration(rand.Intn(4)+1) * time.Second)
		// если произошло прерывание
		case <-jo.quit:
			jo.Done()
			log.Println("#", n, "-- Ok, sir!")
			return
		}
	}

}

func main() {
	var jo JobOffer
	var n int
	//создаем канал для получения сигналов
	sigs := make(chan os.Signal, 1)
	// сигналов о прерывании или о вмешивании
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	if initwork(&jo, &n) == false {
		return
	}

	//зщапуск работников
	for n > 0 {
		go work(&jo, n)
		n--
	}
	tasks := [...]interface{}{"poopin'", "kweelin'", "yahoin'", "selebraitin'", "snifin' around", "smilin'"}
	// тут без комментариев
	for {
		select {
		//ситуация — всё кроме момента получеия данных  в канал с сигналами
		default:
			rand.Seed(time.Now().UnixNano())
			jo.mch <- tasks[rand.Intn(len(tasks))]
		// прерывание или прекращение работы
		case <-sigs:
			//закрываем канал
			close(jo.quit)
			log.Println("-- Stop workin', you fools!")
			// ждем когда все работники перестанут работать
			jo.Wait()
			return
		}
	}
}
