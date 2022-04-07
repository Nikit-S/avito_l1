package main

import (
	"fmt"
)

//попытался сам двигаясь посимвольно, но там еще нужно добавить обработку пустых строк итд
func main() {
	arr1 := []rune(" snow   dog   sun  ")
	arr2 := make([]rune, len(arr1))

	fmt.Println(string(arr1))

	arr2_len := len(arr2)
	l := 0
	for i := range arr1 {

		if (arr1[i] == ' ' && i != 0 && arr1[i-1] != ' ') || i == arr2_len-1 {

			if i != arr2_len-1 {
				copy(arr2[arr2_len-i-1:arr2_len-l], arr1[l:i])
			} else {
				copy(arr2[arr2_len-i-1:arr2_len-l], arr1[l:i+1])
			}
			l = i
		}
	}
	fmt.Printf("|%s|\n", string(arr2))
}
