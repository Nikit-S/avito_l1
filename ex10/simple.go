package main

import "fmt"

func main() {
	m := make(map[int][]float64)
	a := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, 1.0, -1.0, -9.1, 0.0}
	var v int
	for _, f := range a {
		if f < 0 {
			v = int(f - 10)
		} else {
			v = int(f)
		}
		group := v / 10 * 10
		fl := f
		_, ok := m[group]
		if !ok {
			m[group] = []float64{}
		}
		m[group] = append(m[group], fl)
	}
	fmt.Println(m)
}
