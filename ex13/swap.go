package main

import "fmt"

func main() {
	a := 12
	b := 42
	//он там сам меняет под капотом все за счет смены ссылок, даже без выделеия памяти
	fmt.Println(a, b)
	a, b = b, a
	fmt.Println(a, b)
}
