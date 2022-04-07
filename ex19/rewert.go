package main

import "fmt"

func main() {
	//rune type is an alias of int32, так что храним что хотим
	arr := []rune{'♛', 'e', 'l', 10000, 'o'}

	fmt.Println(string(arr))
	l := len(arr) - 1
	for i := 0; i <= l; i++ {
		arr[l], arr[i] = arr[i], arr[l]
		l--
	}
	fmt.Println(string(arr))
	for _, e := range arr {
		fmt.Printf("'%c': %U\n", e, e)
	}

}
