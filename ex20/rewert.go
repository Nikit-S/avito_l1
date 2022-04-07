package main

import (
	"fmt"
	"strings"
)

//а можно не изобретать велосипед и пользоваться встроенными функциями
func main() {
	//разборка на поля игнорируя все символы пробелов табуляций и пр.
	arr := strings.Fields("  sun	 dog snow ")
	left := 0
	right := len(arr) - 1

	//а дальше перевертыш как по буквам
	for left <= right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
	//сборка с красивым сепаратором
	fmt.Println(strings.Join(arr, " "))
}
