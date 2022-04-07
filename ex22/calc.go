package main

import (
	"fmt"
	"math/big"
	"os"
)

//под капотом хранение каждой цифры в массиве и операции с ними, поэтому все это
//не очень быстрое
func main() {
	a := new(big.Int)
	b := new(big.Int)
	res := new(big.Int)
	a.SetString(os.Args[1], 10)
	b.SetString(os.Args[2], 10)

	res.Add(a, b)
	fmt.Println("Add: ", res)

	res.Div(a, b)
	fmt.Println("Div: ", res)

	res.Mul(a, b)
	fmt.Println("Mul: ", res)

	res.Sub(a, b)
	fmt.Println("Sub: ", res)
}
