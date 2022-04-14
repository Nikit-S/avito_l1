package main

import (
	"fmt"
	"reflect"
	"sync"
)

type MyType struct{}

func main() {
	//создаем переменные разных типов
	var wg sync.WaitGroup
	var mut sync.Cond
	var mt MyType
	arr := []interface{}{}
	arr = append(arr, 1, true, "hello", wg, mut, mt)

	//способ один через Тип под копотом у фмт
	for _, e := range arr {
		fmt.Print(e)
		fmt.Printf(" is %T\n", e)
	}

	fmt.Println("")
	// способ два, через библиотеку рефлект (но она может быть медленной)
	for _, e := range arr {
		fmt.Print(e)
		fmt.Print(" is ", reflect.TypeOf(e), "\n")
	}
	fmt.Println(reflect.TypeOf(arr[3]) == reflect.TypeOf(arr[4]))
	fmt.Println(reflect.TypeOf(arr[3]) == reflect.TypeOf(arr[3]))
}

/*
switch v := a.(type) {
// v это тип переменной a
case int:
   fmt.Print("через switch int '", v, "'\n")
case bool:
   fmt.Print("через switch bool '", v, "'\n")
case string:
   fmt.Print("через switch string '", v, "'\n")
case chan struct{}:
   fmt.Print("через switch chan '", v, "'\n")
default:
   fmt.Print("через switch unknown")
}
*/
