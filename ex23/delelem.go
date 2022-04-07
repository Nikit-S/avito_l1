package main

import "fmt"

//не очень быстро потому что двигаем элементы
func remsave(arr []int, i int) []int {
	return append((arr)[0:i], (arr)[i+1:]...)
}

//быстро так как меняем местами два
func remfast(arr []int, i int) []int {
	arr[i] = arr[len(arr)-1]
	return arr[:len(arr)-1]
}

func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7}

	fmt.Println(arr)
	arr = remsave(arr, 3)
	fmt.Println(arr)
	arr = remfast(arr, 2)
	fmt.Println(arr)

}
