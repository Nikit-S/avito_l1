package main

import (
	"fmt"
	"math/rand"
	"time"
)

//элегантное решение
func quicksort(a []int) []int {

	//сразу делаем возврат если в слайсе один элемент == все отсортировано
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivot := rand.Int() % len(a)

	//ставим поворотный элемент в конец
	a[pivot], a[right] = a[right], a[pivot]

	//двигаемся и по ходу при необходимости меняем элементы меньшие с последним
	//запомнившимся элементом который больше
	for i := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	//в конце меняем поворотный на полседний что больше и получаем слайс в котором
	// все меньшие элементы сдева а все большие слева, с ними дальше и работаем
	a[left], a[right] = a[right], a[left]
	quicksort(a[:left])
	quicksort(a[left+1:])

	return a
}

func main() {
	rand.Seed(time.Now().UnixNano())
	//slice := []int{13, -814, 154, 46, -481, 192, 535, -695, -711, -386, -243, 206, -447, 931, 921, -69, -523, 312, 227, 395}
	slice := []int{'A', 'C', 'D', 'k', 'a'}
	//slice := []int{4, 2, 7, 9, 8, -1, -2}
	fmt.Println("\n--- Unsorted --- \n\n", slice)
	quicksort(slice)
	fmt.Println("\n--- Sorted ---\n\n", slice)
}
