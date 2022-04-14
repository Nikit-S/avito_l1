package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type sec int

//универслаьный метод который либо ждет подачи в канал через Т секунд или
//закрытия контекста
func (t sec) canc(ctx context.Context, wg *sync.WaitGroup, str string) {
	wg.Add(1)
	select {
	case <-time.After(time.Duration(t) * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(str, ctx.Err())
	}
	wg.Done()
}

//сигнальное закрытие
func (t sec) quitsig(quit chan os.Signal, wg *sync.WaitGroup, str string) {
	wg.Add(1)
	<-quit
	fmt.Println(str)
	wg.Done()
}

// еще есть переменная переданная по указателю
func main() {
	var wg sync.WaitGroup
	var t sec
	t = 5

	//общий контекст
	ctx := context.Background()
	//контекст с возможностью отмены
	ctxc, cancelc := context.WithCancel(ctx)
	//контекст с таймаутом через две секунды после старта
	ctxt, cancelt := context.WithTimeout(ctx, 2*time.Second)
	//конткекст с окончанием в опеределнный момент
	ctxd, canceld := context.WithDeadline(ctx, time.Now().Add(1*time.Second))

	// вызов функций закрытия контекстов на случай ошибок
	defer canceld()
	defer cancelt()
	//без defer потому что для него cancel это основной способ остановки
	cancelc()
	go t.canc(ctxc, &wg, "Cancel")
	go t.canc(ctxt, &wg, "Cancel Timeout")
	go t.canc(ctxd, &wg, "Cancel Deadline")

	//горутина с завершением по испольнению
	go func() {
		wg.Add(1)
		fmt.Println("just leaving")
		wg.Done()
	}()

	sigs := make(chan os.Signal, 1)
	// сигналов о прерывании или о вмешивании
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go t.quitsig(sigs, &wg, "SigChan")

	//горутина с завершением по закртыию канала
	cch := make(chan struct{})
	go func(ch chan struct{}) {
		wg.Add(1)
		fmt.Println("chan closing wait")
		<-ch
		fmt.Println("chan closing success")
		wg.Done()
	}(cch)
	go func(ch chan struct{}) {
		wg.Add(1)
		time.Sleep(2 * time.Second)
		fmt.Println("chan closing command")
		close(ch)
		wg.Done()
	}(cch)

	//ожидания всех горутин
	wg.Wait()
	fmt.Println("All routines ended")
}
