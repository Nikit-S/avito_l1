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
	arr := make([]interface{}, 0)
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
