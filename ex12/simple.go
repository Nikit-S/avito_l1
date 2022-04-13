package main

import "fmt"

func main() {

	m := make(map[interface{}]interface{})
	fmt.Print("Repeats: ")
	for _, e := range words {

		_, ok := m[e]
		if ok {
			fmt.Print(e, ", ")
		}
		m[e] = struct{}{}
	}
	fmt.Print("\nUnique: ")
	for k := range m {
		fmt.Print(k, ", ")
	}

}
