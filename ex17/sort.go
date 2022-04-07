package main

import (
	"fmt"
)

//стандартный алгоритм в библиотеке сортировки Sort

func Search(n int, f func(int) bool) int {
	// Define f(-1) == false and f(n) == true.
	// Invariant: f(i-1) == false, f(j) == true.
	i, j := 0, n
	for i < j {

		//поиск середины
		// если до этого двигали левую шраницу то получми 3/2  >> 1 = 3/4 от имеющегося массива
		// а если правую то получим 1/2 >> 1 = 1/4
		h := int(uint(i+j) >> 1) // avoid overflow when computing h
		// i ≤ h < j
		if !f(h) {
			// двигаем левую границу до пределов середины
			i = h + 1 // preserves f(i-1) == false
		} else {
			//двигаем праву. границу до середины
			j = h // preserves f(j) == true
		}
	}
	// i == j, f(i-1) == false, and f(j) (= f(i)) == true  =>  answer is i.
	return i
}

func main() {
	a := []int{1, 3, 6, 10, 15, 21, 28, 36, 45, 55}
	x := 6

	i := Search(len(a), func(i int) bool { return a[i] >= x })
	if i < len(a) && a[i] == x {
		fmt.Printf("found %d at index %d in %v\n", x, i, a)
	} else {
		fmt.Printf("%d not found in %v\n", x, a)
	}
}
