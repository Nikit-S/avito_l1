package main

import (
	"log"
	"time"
)

func Sleep(x int) {
	<-time.After(time.Second * time.Duration(x))
}

//https://xwu64.github.io/2019/02/27/Understanding-Golang-sleep-function/
//вот так я не хочу писать
//можно еще и до ассемблера пойти в общем то

func main() {
	log.Println("a")
	Sleep(1)
	log.Println("b")
}
